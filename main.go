package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	// "github.com/nasOS-official/gfb"
)

// var resX, resY int = gfb.GetResolution("fb0")
// var fb []uint8 = gfb.InitFb()

// func logo() {
// 	if (resX >= 640) && (resY >= 480) {
// 		// drawTestRainbow(fb, (resX-resY)/2, resY+((resX-resY)/2), 0, resY)
// 		// gfb.DrawCircle(fb, resY/2, int(float64(resX)*float64(80)/100.0), 128, 255, 255, 255)
// 		x_center := int(float64(resX) * float64(80) / 100.0)
// 		y_center := resY / 2
// 		radius := 128
// 		gfb.DrawCircle(fb, y_center, x_center, radius, 255, 255, 255)

// 		gfb.DrawRectangle(fb, x_center-61, x_center-35, y_center-37, y_center+67, 27, 173, 255)
// 		gfb.DrawRectangle(fb, x_center-13, x_center+13, y_center-64, y_center+67, 27, 173, 255)
// 		gfb.DrawRectangle(fb, x_center+35, x_center+61, y_center-6, y_center+67, 27, 173, 255)

//			gfb.UpdateScreen(fb)
//		}
//	}
const welcomeru = "Добро пожаловать в программу установки SerfOS! \nВыберите язык системы используя стрелки. Нажмите Enter для продолжения."
const welcomeen = "Welcome to the SerfOS Installer! \nselect the system language using the arrows. Press Enter to continue."
const Licenceru = `Лицензионное соглашение!

Здравствуйте, благодарим за выбор SerfOS.
Данный продукт частично* распространяется по лицензии
GNU General Public License 3.0
Более подробно вы можете прочитать здесь:
https://www.gnu.org/licenses/gpl-3.0.html#license-text

*под раздел частично попадает: ядро Linux, графическая среда
Gnome и входящее в дистрибутив Linux программное обеспечение,
за исключением продуктов Crystal Project (калькулятор, блокнот).

Также, перед использованием программных продуктов Crystal,
пожалуйста, ознакомьтесь с условиями настоящего 
лицензионного соглашения.
Пользуясь программными продуктами Crystal,
Вы соглашаетесь с тем, что:
А) Программа имеет закрытый исходный код и вы не имеете право изменять
его или как-либо модифицировать.
Б) Программа предоставляется "как есть", без гарантийных обязательств,
явных или подразумеваемых, либо предусмотренных законодательством, 
включая, но не ограничиваясь этим, гарантии качества,
производительности, пригодности для продажи или для определенной цели.
В) Также не предоставляется никаких гарантий, созданных в результате
заключения сделки, использования или продаж. Crystal Project не 
гарантирует бесперебойную, своевременную и безошибочную работу 
программного обеспечения. ни при каких Crystal Project не несут 
ответственность за ущерб или убытки, вызванные использованием или 
невозможностью использования данного продукта. 
Г) ПО по данному соглашению предоставляется без явных или 
подразумеваемых гарантий о не нарушениях, и разработчик не дает 
гарантий о не нарушениях любых патентов, авторских прав, торговых
секретов или других прав собственности.
Если вы не согласны с условиями лицензии,
немедленно прекратите использование данного ПО!

Авторы дистрибутива:
Игнатьев Илья
Егор aka Linux_Tester
В дистрибутиве использованы обои
Daniel Myslivets https://github.com/myslivets`
const Licenceen = `License Agreement!

Hello, thank you for choosing SerfOS.
This product is partially* distributed under license
GNU General Public License 3.0
You can read more details here:
https://www.gnu.org/licenses/gpl-3.0.html#license-text

*under the section partially falls: Linux kernel, graphical environment
Gnome and the software included in the Linux distribution,
with the exception of Crystal Project products (calculator, notepad).

Also, before using Crystal software products,
please read the terms of this
license agreement.
Using Crystal software products,
You agree that:
A) The Program has closed source code and you have no right to change
it or modify it in any way.
B) The Program is provided "as is", without warranty,
express or implied, or statutory,
including, but not limited to, warranties of quality,
performance, merchantability or fitness for a particular purpose.
C) There are also no guarantees created as a result
of the conclusion of the transaction, use or sales. Crystal Project is not 
guarantees uninterrupted, timely and error-free operation
of the software. Under no circumstances will Crystal Project be
liable for damages or losses caused by the use or
inability to use this product. 
D) The software under this agreement is provided without express or
implied warranties of non-infringement, and the developer does not
guarantee non-infringement of any patents, copyrights, trade
secrets or other proprietary rights.
If you do not agree to the terms of the license,
immediately stop using this software!

Distribution authors:
Ignatiev Ilya
Egor aka Linux_Tester
Wallpaper is used in the distribution
Daniel Myslivets https://github.com/myslivets`

func showmenu(elem int, menu []string, title string) {
	fmt.Printf("\033c")
	// logo()
	fmt.Printf("\x1b[36m" + title + "\x1b[0m\n")
	for i := 0; i < len(menu); i++ {
		if elem == i {
			fmt.Printf("\x1b[47;30m" + menu[i] + "\x1b[0m\n")
		} else {
			fmt.Println(menu[i])
		}
	}
}
func selectlang(language string) string {
	// logo()
	title := welcomeru
	item := 0
	menu := []string{"Русский", "English"}
	showmenu(item, menu, title)
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {

		switch key.Code {
		case keys.Up:
			if item != 0 {
				item--
			} else {
				item = len(menu) - 1
			}
			if item == 0 {
				title = welcomeru
				language = "ru"

			} else {
				title = welcomeen
				language = "en"
			}

			showmenu(item, menu, title)

		case keys.Down:

			if item != len(menu)-1 {
				item++
			} else {
				item = 0
			}
			if item == 0 {
				title = welcomeru
				language = "ru"
			} else {
				title = welcomeen
				language = "en"
			}
			showmenu(item, menu, title)
		case keys.Enter:
			switch item {
			case 0:
				return true, nil
			case 1:
				return true, nil
			}
		case keys.CtrlC:
			return true, nil // Stop listener by returning true on Ctrl+C
		}

		return false, nil // Return false to continue listening
	})
	return language
}
func showLicense(language string) {
	fmt.Printf("\033c")
	// logo()
	title := ""
	exit := 0
	menu := []string{"Do not accept", "Accept"}
	if language == "ru" {
		title = Licenceru
		menu = []string{"Не принимаю", "Принимаю"}

	} else {
		title = Licenceen
		menu = []string{"Do not accept", "Accept"}
	}
	item := 0

	showmenu(item, menu, title)
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {

		switch key.Code {
		case keys.Up:
			if item != 0 {
				item--
			} else {
				item = len(menu) - 1
			}

			showmenu(item, menu, title)

		case keys.Down:

			if item != len(menu)-1 {
				item++
			} else {
				item = 0
			}

			showmenu(item, menu, title)
		case keys.Enter:
			switch item {
			case 0:
				exit = 1
				return true, nil
			case 1:
				exit = 0
				return true, nil
			}
		case keys.CtrlC:
			return true, nil // Stop listener by returning true on Ctrl+C
		}

		return false, nil // Return false to continue listening
	})
	if exit == 1 {
		cmd := exec.Command("shutdown", "now")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		// Запускаем команду
		_ = cmd.Start()
		_ = cmd.Wait()
		os.Exit(0)
	}
}

func partiton(language string) string {
	title := ""
	drive := ""
	if language == "ru" {
		title = "Выберите диск для разметки"

	} else {
		title = "Select disk drive for partitioning"
	}
	_ = title
	devices, _ := filepath.Glob("/dev/[sS]d[a-zA-Z]")

	nvmeDevices, _ := filepath.Glob("/dev/nvme[0-9]n[0-9]")

	mmcDevices, _ := filepath.Glob("/dev/mmcblk[0-9]")

	allDevices := append(devices, nvmeDevices...)
	allDevices = append(allDevices, mmcDevices...)
	menu := allDevices
	item := 0

	showmenu(item, menu, title)
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {

		switch key.Code {
		case keys.Up:
			if item != 0 {
				item--
			} else {
				item = len(menu) - 1
			}

			showmenu(item, menu, title)

		case keys.Down:

			if item != len(menu)-1 {
				item++
			} else {
				item = 0
			}

			showmenu(item, menu, title)
		case keys.Enter:
			drive = allDevices[item]
			args := strings.Split("cfdisk"+" "+drive, " ")
			cmd := exec.Command(args[0], args[1])

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin

			// Запускаем команду
			err = cmd.Start()
			if err != nil {
				log.Fatal(err)
			}
			_ = cmd.Wait()
			return true, nil
		case keys.CtrlC:
			return true, nil // Stop listener by returning true on Ctrl+C
		}

		return false, nil // Return false to continue listening
	})
	return drive

}

func partboot(language string, instdrive string) string {
	devices := []string{}
	drive := ""
	if _, err := os.Stat("/sys/firmware/efi"); err == nil {
		// Объявляем пустой срез перед условным оператором
		re := regexp.MustCompile(`/dev/[sS]d[a-zA-Z]`)

		if re.MatchString(instdrive) {

			devices, _ = filepath.Glob(instdrive + "[0-9]") // Присваиваем значение внутри блока кода if
		} else {
			devices, _ = filepath.Glob(instdrive + "p[0-9]") // Присваиваем значение внутри блока кода else
		}

		title := ""
		if language == "ru" {
			title = "Выберите раздел для загрузчика (fat32)"

		} else {
			title = "Select partition for boot loader (fat32)"
		}
		menu := devices
		item := 0

		showmenu(item, menu, title)
		keyboard.Listen(func(key keys.Key) (stop bool, err error) {

			switch key.Code {
			case keys.Up:
				if item != 0 {
					item--
				} else {
					item = len(menu) - 1
				}

				showmenu(item, menu, title)

			case keys.Down:

				if item != len(menu)-1 {
					item++
				} else {
					item = 0
				}

				showmenu(item, menu, title)
			case keys.Enter:
				drive = menu[item]
				return true, nil
			case keys.CtrlC:
				return true, nil // Stop listener by returning true on Ctrl+C
			}

			return false, nil // Return false to continue listening
		})

	}
	return drive
}

func partsel(language string, instdrive string) string {
	devices := []string{} // Объявляем пустой срез перед условным оператором
	re := regexp.MustCompile(`/dev/[sS]d[a-zA-Z]`)

	if re.MatchString(instdrive) {

		devices, _ = filepath.Glob(instdrive + "[0-9]") // Присваиваем значение внутри блока кода if
	} else {
		devices, _ = filepath.Glob(instdrive + "p[0-9]") // Присваиваем значение внутри блока кода else
	}

	title := ""
	if language == "ru" {
		title = "Выберите раздел для установки системы"

	} else {
		title = "Select partition for installation system"
	}
	drive := ""
	menu := devices
	item := 0

	showmenu(item, menu, title)
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {

		switch key.Code {
		case keys.Up:
			if item != 0 {
				item--
			} else {
				item = len(menu) - 1
			}

			showmenu(item, menu, title)

		case keys.Down:

			if item != len(menu)-1 {
				item++
			} else {
				item = 0
			}

			showmenu(item, menu, title)
		case keys.Enter:
			drive = menu[item]
			return true, nil
		case keys.CtrlC:
			return true, nil // Stop listener by returning true on Ctrl+C
		}

		return false, nil // Return false to continue listening
	})
	return drive

}
func sysinstall(language string, instpart string, bootpart string, instdrive string) {
	title := ""
	if language == "ru" {
		title = "Идет установка, пожалуйста, подождите. Это может занять несколько минут."
	} else {
		title = "Installing, please wait. This may take a few minutes."
	}
	args := strings.Split("mkfs.ext4"+" "+"-q"+" "+instpart, " ")
	cmd := exec.Command(args[0], args[1], args[2])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Запускаем команду
	_ = cmd.Start()
	_ = cmd.Wait()
	args = strings.Split("mount"+" "+instpart+" "+"/mnt", " ")
	cmd = exec.Command(args[0], args[1], args[2])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Запускаем команду
	_ = cmd.Start()
	_ = cmd.Wait()
	args = strings.Split("tar xvzf system.tar.gz -C /mnt", " ")
	cmd = exec.Command(args[0], args[1], args[2], args[3], args[4])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Запускаем команду
	_ = cmd.Start()
	_ = cmd.Wait()

	if bootpart != "" {

		fmt.Println(title)

		args := strings.Split("mkdir -p /mnt/boot/efi", " ")
		cmd := exec.Command(args[0], args[1], args[2])
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		// Запускаем команду
		_ = cmd.Start()
		_ = cmd.Wait()
		args = strings.Split("mkfs.fat -F32 "+bootpart, " ")
		cmd = exec.Command(args[0], args[1], args[2])
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		// Запускаем команду
		_ = cmd.Start()
		_ = cmd.Wait()
		args = strings.Split("mount"+" "+bootpart+" "+"/mnt/boot/efi", " ")
		cmd = exec.Command(args[0], args[1], args[2])
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		// Запускаем команду
		_ = cmd.Start()
		_ = cmd.Wait()
		cmd = exec.Command("bash", "./tab.sh")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		// Запускаем команду
		_ = cmd.Start()
		_ = cmd.Wait()

		cmd = exec.Command("arch-chroot", "/mnt", "grub-install")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		// Запускаем команду
		_ = cmd.Start()
		_ = cmd.Wait()

	} else {
		cmd = exec.Command("bash", "./tab.sh")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		// Запускаем команду
		_ = cmd.Start()
		_ = cmd.Wait()
		args = strings.Split("arch-chroot /mnt grub-install "+instdrive, " ")
		cmd = exec.Command(args[0], args[1], args[2], args[3])
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		// Запускаем команду
		_ = cmd.Start()
		_ = cmd.Wait()

	}
	args = strings.Split("arch-chroot /mnt update-grub ", " ")
	cmd = exec.Command(args[0], args[1], args[2])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Запускаем команду
	_ = cmd.Start()
	_ = cmd.Wait()

	fmt.Printf("\033c")
	if language == "ru" {
		fmt.Println("Пожалуйста, введите пароль суперпользователя. (Пароль не отображается)")
	} else {
		fmt.Println("Please enter the superuser password. (Password is not displayed)")
	}
	args = strings.Split("arch-chroot /mnt passwd", " ")
	cmd = exec.Command(args[0], args[1], args[2])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Запускаем команду
	_ = cmd.Start()
	_ = cmd.Wait()
	if language == "ru" {
		fmt.Println("Пожалуйста, введите пароль пользователя. (Пароль не отображается)")
	} else {
		fmt.Println("Please enter the user password. (Password is not displayed)")
	}
	args = strings.Split("arch-chroot /mnt passwd live", " ")
	cmd = exec.Command(args[0], args[1], args[2], args[3])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Запускаем команду
	_ = cmd.Start()
	_ = cmd.Wait()
	cmd = exec.Command("reboot")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Запускаем команду
	_ = cmd.Start()
	_ = cmd.Wait()

}
func main() {
	language := "ru"
	language = selectlang(language)
	showLicense(language)
	instdrive := partiton(language)
	fmt.Println(instdrive)
	bootpart := partboot(language, instdrive)
	instpart := partsel(language, instdrive)
	sysinstall(language, instpart, bootpart, instdrive)
	// selectpart(language)
	os.Exit(0)
}

//юзеры
//установка
