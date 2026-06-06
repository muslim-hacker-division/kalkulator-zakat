package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.DarkTheme())
	
	myWindow := myApp.NewWindow("Zakat Hub - Responsif")
	myWindow.Resize(fyne.NewSize(360, 500))

	// ==========================================
	// TAB 1: ZAKAT MAL
	// ==========================================
	labelJudulMal := widget.NewLabelWithStyle("KALKULATOR ZAKAT MAL", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	
	inputHarta := widget.NewEntry()
	inputHarta.SetPlaceHolder("Contoh: 120000000")

	labelHasil := widget.NewRichTextFromMarkdown("")
	labelHasil.Wrapping = fyne.TextWrapWord 

	tombolHitung := widget.NewButton("Hitung Zakat Sekarang", func() {
		hartaStr := inputHarta.Text
		hartaFl, err := strconv.ParseFloat(hartaStr, 64)

		if err != nil {
			labelHasil.ParseMarkdown("**Status:** `Error: Input salah!`")
			return
		}

		hargaEmas := 1400000.0
		nishab := 85.0 * hargaEmas

		if hartaFl >= nishab {
			zakat := hartaFl * 0.025
			labelHasil.ParseMarkdown(fmt.Sprintf("### STATUS: WAJIB ZAKAT\n**Zakat:** Rp %.0f", zakat))
		} else {
			labelHasil.ParseMarkdown("### STATUS: BELUM WAJIB\nHarta belum mencapai batas Nishab.")
		}
	})

	formKalkulator := widget.NewForm(
		widget.NewFormItem("Total Harta", inputHarta),
	)

	dalilMal := widget.NewRichTextFromMarkdown(
		"### Dalil Zakat Mal:\n" +
			"> *\"Ambil-lah zakat dari sebagian harta mereka...\"* (QS. At-Taubah: 103)\n\n" +
			"**Syarat:** Halal, milik penuh, mencapai nishab, & haul 1 tahun.",
	)
	dalilMal.Wrapping = fyne.TextWrapWord 

	boxMal := container.NewVBox(
		labelJudulMal,
		widget.NewSeparator(),
		formKalkulator,
		tombolHitung,
		widget.NewSeparator(),
		labelHasil,
		widget.NewSeparator(),
		dalilMal,
	)
	
	kontenZakatMal := container.NewScroll(boxMal)

	// ==========================================
	// TAB 2: ZAKAT FITRAH
	// ==========================================
	edukasiFitrah := widget.NewRichTextFromMarkdown(
		"# ZAKAT FITRAH\n***\n" +
			"**Dalil:**\n" +
			"> *\"Rasulullah SAW mewajibkan zakat fitrah sebagai penyuci...\"* (HR. Abu Dawud)\n\n" +
			"**Takaran:** 2,5 kg / 3,5 liter beras per jiwa.\n\n" +
			"**Waktu:**\n" +
			"* *Afdhal:* Pagi Idul Fitri sebelum shalat.\n" +
			"* *Wajib:* Matahari terbenam akhir Ramadhan.\n" +
			"* *Haram:* Setelah shalat Id selesai.",
	)
	edukasiFitrah.Wrapping = fyne.TextWrapWord
	kontenZakatFitrah := container.NewScroll(container.NewVBox(edukasiFitrah))

	// ==========================================
	// TAB 3: SEDEKAH
	// ==========================================
	edukasiSedekah := widget.NewRichTextFromMarkdown(
		"# INFAK & SEDEKAH\n***\n" +
			"**Keutamaan:**\n" +
			"> *\"Perumpamaan orang yang menginfakkan hartanya...\"* (QS. Al-Baqarah: 261)\n\n" +
			"**Sedekah Non-Materi:**\n" +
			"1. Menyingkirkan duri di jalan.\n" +
			"2. Menolong dengan tenaga/ilmu.\n" +
			"3. Tersenyum tulus.",
	)
	edukasiSedekah.Wrapping = fyne.TextWrapWord
	kontenSedekah := container.NewScroll(container.NewVBox(edukasiSedekah))

	// ==========================================
	// MENU TAB
	// ==========================================
	menuTabs := container.NewAppTabs(
		container.NewTabItem("Zakat Mal", kontenZakatMal),
		container.NewTabItem("Zakat Fitrah", kontenZakatFitrah),
		container.NewTabItem("Sedekah", kontenSedekah),
	)
	menuTabs.SetTabLocation(container.TabLocationTop)

	myWindow.SetContent(menuTabs)
	myWindow.ShowAndRun()
}
