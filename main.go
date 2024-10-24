package main

import (
	"fmt"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"clod_Unauthorized_tool/internal/docker"
	"clod_Unauthorized_tool/internal/k8sapiserver"
	"clod_Unauthorized_tool/internal/k8sdashboard"
	"clod_Unauthorized_tool/internal/etcd"
)

func main() {
	a := app.New()
	w := a.NewWindow("Clod Unauthorized Tool")
	w.Resize(fyne.NewSize(600, 400))

	urlEntry := widget.NewEntry()
	urlEntry.SetPlaceHolder("Enter URL")

	fileEntry := widget.NewEntry()
	fileEntry.SetPlaceHolder("Enter file path")

	selectFileButton := widget.NewButton("Select File", func() {
		fileDialog := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
			if err == nil && uc != nil {
				fileEntry.SetText(uc.URI().Path())
				uc.Close()
			}
		}, w)
		fileDialog.Show()
	})

	moduleSelect := widget.NewSelect([]string{"docker", "k8sapiserver", "k8sdashboard", "etcd", "all"}, nil)
	moduleSelect.SetSelected("all")

	var results []string
	var selectedIndex int

	vulnerableURLs := widget.NewList(
		func() int { return len(results) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(results[i])
		},
	)

	vulnerableURLs.OnSelected = func(id widget.ListItemID) {
		selectedIndex = id
	}

	checkButton := widget.NewButton("Check", func() {
		results = []string{}
		vulnerableURLs.Refresh()

		urls := []string{}
		if urlEntry.Text != "" {
			urls = append(urls, urlEntry.Text)
		}
		if fileEntry.Text != "" {
			data, err := os.ReadFile(fileEntry.Text)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
			urls = append(urls, strings.Split(string(data), "\n")...)
		}

		for _, url := range urls {
			if moduleSelect.Selected == "all" || moduleSelect.Selected == "docker" {
				if docker.Check(url) {
					results = append(results, fmt.Sprintf("Docker: %s", url))
				} else {
					results = append(results, fmt.Sprintf("No Docker vulnerability detected: %s", url))
				}
				vulnerableURLs.Refresh()
			}
			if moduleSelect.Selected == "all" || moduleSelect.Selected == "k8sapiserver" {
				if k8sapiserver.Check(url) {
					results = append(results, fmt.Sprintf("Kubernetes API Server: %s", url))
				} else {
					results = append(results, fmt.Sprintf("No Kubernetes API Server vulnerability detected: %s", url))
				}
				vulnerableURLs.Refresh()
			}
			if moduleSelect.Selected == "all" || moduleSelect.Selected == "k8sdashboard" {
				if k8sdashboard.Check(url) {
					results = append(results, fmt.Sprintf("Kubernetes Dashboard: %s", url))
				} else {
					results = append(results, fmt.Sprintf("No Kubernetes Dashboard vulnerability detected: %s", url))
				}
				vulnerableURLs.Refresh()
			}
			if moduleSelect.Selected == "all" || moduleSelect.Selected == "etcd" {
				if etcd.Check(url) {
					results = append(results, fmt.Sprintf("etcd: %s", url))
				} else {
					results = append(results, fmt.Sprintf("No etcd vulnerability detected: %s", url))
				}
				vulnerableURLs.Refresh()
			}
		}
	})

	exploitButton := widget.NewButton("Exploit", func() {
		if len(results) > 0 && selectedIndex >= 0 && selectedIndex < len(results) {
			selectedResult := results[selectedIndex]
			parts := strings.SplitN(selectedResult, ": ", 2)
			if len(parts) == 2 {
				moduleType := parts[0]
				selectedURL := parts[1]
				switch moduleType {
				case "Docker":
					docker.Exploit(selectedURL)
				case "Kubernetes API Server":
					k8sapiserver.Exploit(selectedURL)
				case "Kubernetes Dashboard":
					k8sdashboard.Exploit(selectedURL)
				case "etcd":
					etcd.Exploit(selectedURL)
				}
			}
		}
	})

	saveResultsButton := widget.NewButton("Save Results", func() {
		if len(results) > 0 {
			saveFileDialog := dialog.NewFileSave(
				func(uc fyne.URIWriteCloser, err error) {
					if err == nil && uc != nil {
						defer uc.Close()
						uc.Write([]byte(strings.Join(results, "\n")))
					}
				}, w)
			saveFileDialog.SetFileName("results.txt")
			saveFileDialog.Show()
		}
	})

	exportVulnerableURLsButton := widget.NewButton("Export Vulnerable URLs", func() {
		if len(results) > 0 {
			saveFileDialog := dialog.NewFileSave(
				func(uc fyne.URIWriteCloser, err error) {
					if err == nil && uc != nil {
						defer uc.Close()
						uc.Write([]byte(strings.Join(results, "\n")))
					}
				}, w)
			saveFileDialog.SetFileName("vulnerable_urls.txt")
			saveFileDialog.Show()
		}
	})

	inputContainer := container.NewVBox(
		widget.NewLabel("URL:"),
		urlEntry,
		widget.NewLabel("File Path:"),
		fileEntry,
		selectFileButton,
		widget.NewLabel("Module:"),
		moduleSelect,
	)

	buttonContainer := container.NewHBox(
		checkButton,
		exploitButton,
		saveResultsButton,
		exportVulnerableURLsButton,
	)

	mainContainer := container.New(layout.NewVBoxLayout(),
		inputContainer,
		buttonContainer,
		widget.NewLabel("Vulnerable URLs:"),
		vulnerableURLs,
	)

	w.SetContent(mainContainer)
	w.ShowAndRun()
}