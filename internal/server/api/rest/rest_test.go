package resthandler

import (
	"best-goph-keeper/internal/client/service/encryption"
	"best-goph-keeper/internal/client/service/randomizer"
	"best-goph-keeper/internal/server/api"
	grpcHandler "best-goph-keeper/internal/server/api/grpc"
	serverConfig "best-goph-keeper/internal/server/config"
	"best-goph-keeper/internal/server/database"
	grpcKeeper "best-goph-keeper/internal/server/proto"
	"best-goph-keeper/internal/server/storage"
	"best-goph-keeper/internal/server/storage/repositories/entity"
	"best-goph-keeper/internal/server/storage/repositories/file"
	"best-goph-keeper/internal/server/storage/repositories/token"
	"best-goph-keeper/internal/server/storage/repositories/user"
	"bytes"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"net/http"
	"testing"
	"time"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func TestRest(t *testing.T) {
	// -- SETUP --
	// initiate postgres container

	container, err := postgres.RunContainer(context.Background(),
		testcontainers.WithImage("docker.io/postgres:latest"),
		postgres.WithDatabase("postgres"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		t.Fatalf("Test containers failed: %v", err)
	}

	err = container.Start(context.Background())
	if err != nil {
		t.Fatalf("Test containers failed: %v", err)
	}

	databaseURI, err := container.ConnectionString(context.Background(), "sslmode=disable")
	if err != nil {
		t.Fatalf("container connection failed: %v", err)
	}

	logger := logrus.New()

	db, err := database.New(&serverConfig.Config{DSN: databaseURI}, logger)
	if err != nil {
		t.Fatalf("Db init failed: %v", err)
	}

	err = db.CreateTablesMigration("file://../../../../migrations")
	if err != nil {
		t.Fatalf("Migration failed: %v", err)
	}

	serverCnfg := &serverConfig.Config{
		AddressGRPC:         "localhost:8080",
		AddressREST:         "localhost:8088",
		DSN:                 databaseURI,
		AccessTokenLifetime: 300 * time.Second,
		DebugLevel:          logrus.DebugLevel,
		FileFolder:          "../../../../data/test_keeper",
		TemplatePathToken:   "../../../../web/templates/token_list.html",
		TemplatePathUser:    "../../../../web/templates/user_list.html",
	}

	// repositories
	userRepository := user.New(db)
	fileRepository := file.New(db)
	storage := storage.New("/tmp")
	entityRepository := entity.New(db)
	tokenRepository := token.New(db)

	// setup server service
	handlerGrpc := *grpcHandler.NewHandler(db, serverCnfg, userRepository, fileRepository, &storage,
		entityRepository, tokenRepository, logger)

	handlerRest := NewHandler(db, serverCnfg, userRepository, tokenRepository, logger)
	routerService := Route(handlerRest)
	rs := chi.NewRouter()
	rs.Mount("/", routerService)

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	t.Run("registration gophkeeper-server", func(t *testing.T) {
		grpcKeeper.RegisterGophkeeperServer(s, &handlerGrpc)
		go api.StartRESTService(rs, serverCnfg, logger)

		go func() {
			if err := s.Serve(lis); err != nil {
				t.Errorf("Server exited with error: %v", err)
				return
			}
		}()
	})

	var authenticatedUser *grpcKeeper.AuthenticationResponse
	username := randomizer.RandStringRunes(10)
	password, _ := encryption.HashPassword("Password-00")

	t.Run("registration", func(t *testing.T) {
		regisResp, err := handlerGrpc.Registration(context.Background(),
			&grpcKeeper.RegistrationRequest{Username: username, Password: password})
		fmt.Println(regisResp)
		assert.NoError(t, err, "registration failed")
	})

	t.Run("authentication", func(t *testing.T) {
		authenticatedUser, err = handlerGrpc.Authentication(context.Background(),
			&grpcKeeper.AuthenticationRequest{Username: username, Password: password})
		assert.NoError(t, err, "authentication failed")
	})

	var b bytes.Buffer

	if _, err := http.Get("http://" + serverCnfg.AddressREST + "/api/"); err != nil {
		t.Errorf("http.Post : %v", err)
	}

	if _, err := http.Post("http://"+serverCnfg.AddressREST+"/api/user/block?username="+username, "", &b); err != nil {
		t.Errorf("http.Post : %v", err)
	}

	if _, err := http.Post("http://"+serverCnfg.AddressREST+"/api/user/block?username=", "", &b); err != nil {
		t.Errorf("http.Post : %v", err)
	}

	if _, err := http.Post("http://"+serverCnfg.AddressREST+"/api/user/unblock?username="+username, "", &b); err != nil {
		t.Errorf("http.Post : %v", err)
	}

	if _, err := http.Post("http://"+serverCnfg.AddressREST+"/api/user/unblock?username=", "", &b); err != nil {
		t.Errorf("http.Post : %v", err)
	}

	if _, err := http.Post("http://"+serverCnfg.AddressREST+"/api/token/block?access_token="+authenticatedUser.AccessToken.Token, "", &b); err != nil {
		t.Errorf("http.Post : %v", err)
	}

	if _, err := http.Post("http://"+serverCnfg.AddressREST+"/api/token/block?access_token=", "", &b); err != nil {
		t.Errorf("http.Post : %v", err)
	}

	if _, err := http.Get("http://" + serverCnfg.AddressREST + "/api/token/" + username); err != nil {
		t.Errorf("http.Post : %v", err)
	}

	stopTime := time.Second
	err = container.Stop(context.Background(), &stopTime)
	if err != nil {
		t.Fatalf("container stop failed: %v", err)
	}
}
