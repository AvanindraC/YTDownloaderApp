package main

import (
	"io"
	"os"
	"os/user"
	"path"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kkdai/youtube/v2"
)

func main() {
	app := app.New()
	//app.Settings().SetTheme(theme.DarkTheme())
	w := app.NewWindow("Youtube Downloader")
	w.SetContent(widget.NewLabel("Download your favourite youtube videos"))
	urll := widget.NewEntry()
	urll.SetPlaceHolder("Enter the video url...")
	filename := widget.NewEntry()
	filename.SetPlaceHolder("Enter the name to save the video/audio...")
	btn := widget.NewButton("Download MP3", func() {})
	btn.Resize(fyne.NewSize(30, 10))
	btn2 := widget.NewButton("Download MP4", func() {
		VidID := urll.Text[17 : len(urll.Text)-1]
		videoID := VidID
		client := youtube.Client{}
		video, err := client.GetVideo(videoID)
		if err != nil {
			panic(err)
		}
		formats := video.Formats.WithAudioChannels()
		stream, _, err := client.GetStream(video, &formats[0])
		if err != nil {
			panic(err)
		}
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		homeDirectory := user.HomeDir
		sd := path.Join(homeDirectory, "YtDown")
		vidname := sd + "/" + filename.Text + ".mp4"
		file, err := os.Create(vidname)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		_, err = io.Copy(file, stream)
		if err != nil {
			panic(err)
		}

	})
	btn2.Resize(fyne.NewSize(30, 10))
	content := container.NewVBox(urll, filename, btn, btn2)

	w.SetContent(content)

	//w.Resize(fyne.NewSize(700, 400))
	w.ShowAndRun()

}
