package main

import (
	"best-goph-keeper/internal/api/client/cmp"
	cleientModel "best-goph-keeper/internal/api/client/model"
	"best-goph-keeper/internal/api/client/service"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"time"
)

func main() {
	application := app.New()
	application.Settings().SetTheme(theme.LightTheme())
	window := application.NewWindow("GophKeeper")
	window.Resize(fyne.NewSize(250, 80))

	//variables
	var dataTblText = [][]string{{"NAME", "DATA", "DESCRIPTION", "CREATED_AT", "UPDATED_AT"}}
	var dataTblCard = [][]string{{"NAME", "PAYMENT SYSTEM", "NUMBER", "HOLDER", "CVC", "END DATE", "CREATED_AT", "UPDATED_AT"}}
	var radioOptions = []string{"Login", "Registration"}
	var user = cleientModel.User{}
	var exist bool
	var valid bool
	var layout string
	layout = "01/02/2006 15:04:05"

	//containers
	var containerRadio *fyne.Container
	var containerFormLogin *fyne.Container
	var containerFormRegistration *fyne.Container
	var containerFormText *fyne.Container
	var containerFormCard *fyne.Container

	//buttons
	var buttonAuth *widget.Button
	var buttonTop *widget.Button
	var buttonText *widget.Button
	var buttonCard *widget.Button
	var buttonTextAdd *widget.Button
	var buttonCardAdd *widget.Button

	//tabs
	var containerTabs *container.AppTabs
	var tblText *widget.Table
	var tblCard *widget.Table
	var tabText *container.TabItem
	var tabCard *container.TabItem

	//entries init
	separator := widget.NewSeparator()
	usernameLoginEntry := widget.NewEntry()
	passwordLoginEntry := widget.NewPasswordEntry()
	usernameRegistrationEntry := widget.NewEntry()
	passwordRegistrationEntry := widget.NewPasswordEntry()
	passwordConfirmationRegistrationEntry := widget.NewPasswordEntry()
	textNameEntry := widget.NewEntry()
	textEntry := widget.NewEntry()
	textDescriptionEntry := widget.NewEntry()
	cardNameEntry := widget.NewEntry()
	paymentSystemEntry := widget.NewEntry()
	numberEntry := widget.NewEntry()
	holderEntry := widget.NewEntry()
	endDateEntry := widget.NewEntry()
	cvcEntry := widget.NewEntry()

	//labels init
	labelAlertAuth := widget.NewLabel("")
	labelAlertText := widget.NewLabel("")
	labelAlertCard := widget.NewLabel("")
	labelAlertAuth.Hide()
	labelAlertText.Hide()
	labelAlertCard.Hide()

	//forms init
	formLogin := cmp.GetFormLogin(usernameLoginEntry, passwordLoginEntry)
	formRegistration := cmp.GetFormRegistration(usernameRegistrationEntry, passwordRegistrationEntry, passwordConfirmationRegistrationEntry)
	formText := cmp.GetFormText(textNameEntry, textEntry, textDescriptionEntry)
	formCard := cmp.GetFormCard(cardNameEntry, paymentSystemEntry, numberEntry, holderEntry, endDateEntry, cvcEntry)

	//radio event
	radioAuth := widget.NewRadioGroup(radioOptions, func(value string) {
		log.Println("Radio set to ", value)
		if value == "Login" {
			window.SetContent(containerFormLogin)
			window.Resize(fyne.NewSize(500, 100))
			window.Show()
		}
		if value == "Registration" {
			window.SetContent(containerFormRegistration)
			window.Resize(fyne.NewSize(500, 100))
			window.Show()
		}
	})

	//buttons event
	buttonTop = widget.NewButton("Обновить данные", func() {
		dataTblText, dataTblCard = service.Sync(user.ID)
		tblText.Resize(fyne.NewSize(float32(len(dataTblText)), float32(len(dataTblText[0]))))
		tblText.Refresh()
		tblCard.Resize(fyne.NewSize(float32(len(dataTblCard)), float32(len(dataTblCard[0]))))
		tblCard.Refresh()
		window.SetContent(containerTabs)
	})
	buttonText = widget.NewButton("Добавить текстовые данные", func() {
		window.SetContent(containerFormText)
		window.Show()
	})
	buttonCard = widget.NewButton("Добавить банковскую карту", func() {
		window.SetContent(containerFormCard)
		window.Show()
	})
	//table text init
	tblText = widget.NewTable(
		func() (int, int) {
			return len(dataTblText), len(dataTblText[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(dataTblText[i.Row][i.Col])
		})
	service.SetDefaultColumnsWidthText(tblText)
	//table card init
	tblCard = widget.NewTable(
		func() (int, int) {
			return len(dataTblCard), len(dataTblCard[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(dataTblCard[i.Row][i.Col])
		})
	service.SetDefaultColumnsWidthCard(tblCard)
	//containerTabs
	tabText = cmp.GetTabTexts(tblText, buttonTop, buttonText)
	tabCard = cmp.GetTabCards(tblCard, buttonTop, buttonCard)
	containerTabs = container.NewAppTabs(tabText, tabCard)
	//auth event
	buttonAuth = widget.NewButton("Submit", func() {
		labelAlertAuth.Show()
		valid = false
		if radioAuth.Selected == "Login" {
			valid = service.ValidateLogin(usernameLoginEntry, passwordLoginEntry, labelAlertAuth)
			if valid {
				user, exist = service.Authentication(usernameLoginEntry.Text, passwordLoginEntry.Text) //ищем в бд
				if exist {
					dataTblText, dataTblCard = service.Sync(user.ID)
					window.SetContent(containerTabs)
					window.Resize(fyne.NewSize(1250, 300))
					window.Show()
				}
			}
		}
		if radioAuth.Selected == "Registration" {
			valid = service.ValidateRegistration(usernameRegistrationEntry, passwordRegistrationEntry, passwordConfirmationRegistrationEntry, labelAlertAuth)
			if valid {
				exist = service.UserExist(usernameRegistrationEntry.Text) //ищем в бд
				if !exist {
					user = service.Registration(usernameRegistrationEntry.Text, passwordRegistrationEntry.Text)
					window.SetContent(containerTabs)
					window.Resize(fyne.NewSize(1250, 300))
					window.Show()
				}
			}
		}
	})
	//text event
	buttonTextAdd = widget.NewButton("Добавить", func() {
		labelAlertText.Show()
		valid = false
		exist = service.SearchByColumn(dataTblText, 0, textNameEntry.Text) //ищем в мапке
		valid = service.ValidateText(exist, textNameEntry, textEntry, textDescriptionEntry, labelAlertText)
		if valid {
			dataTblText = append(dataTblText, []string{textNameEntry.Text, textEntry.Text, textDescriptionEntry.Text,
				time.Now().Format(layout), time.Now().Format(layout)})

			service.ClearText(textNameEntry, textEntry, textDescriptionEntry)
			log.Print("Текст добавлен")

			labelAlertText.Hide()
			formText.Refresh()
			window.SetContent(containerTabs)
			window.Show()
		}
		log.Print(dataTblText)
	})
	//card event
	buttonCardAdd = widget.NewButton("Добавить", func() {
		labelAlertCard.Show()
		valid = false
		exist = service.SearchByColumn(dataTblCard, 0, cardNameEntry.Text) //ищем в мапке
		valid = service.ValidateCard(exist, cardNameEntry, paymentSystemEntry, numberEntry, holderEntry, endDateEntry, cvcEntry, labelAlertCard)
		if valid {
			layout := "01/02/2006 15:04:05"
			dataTblCard = append(dataTblCard, []string{cardNameEntry.Text, paymentSystemEntry.Text, numberEntry.Text, holderEntry.Text,
				cvcEntry.Text, endDateEntry.Text, time.Now().Format(layout), time.Now().Format(layout)})

			service.ClearCard(cardNameEntry, paymentSystemEntry, numberEntry, holderEntry, endDateEntry, cvcEntry)
			log.Print("Карта добавлена")

			labelAlertCard.Hide()
			formCard.Refresh()
			window.SetContent(containerTabs)
			window.Show()
		}
		log.Print(dataTblCard)
	})
	//containers init
	containerRadio = container.NewVBox(radioAuth)
	containerFormLogin = container.NewVBox(formLogin, buttonAuth, labelAlertAuth, separator, radioAuth)
	containerFormRegistration = container.NewVBox(formRegistration, buttonAuth, labelAlertAuth, separator, radioAuth)
	containerFormText = container.NewVBox(formText, buttonTextAdd, labelAlertText)
	containerFormCard = container.NewVBox(formCard, buttonCardAdd, labelAlertCard)

	window.SetContent(containerRadio)
	window.ShowAndRun()
}
