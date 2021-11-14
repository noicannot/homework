package pkg

import "fmt"

// Logo print banner to console ...
func Logo()  {
	// Generate with `ansi-shadow` font at `https://www.bootschool.net/ascii`
	banner :=
`██████╗  ██████╗  ██████╗ ██╗  ██╗     █████╗ ██████╗ ██╗
██╔══██╗██╔═══██╗██╔═══██╗██║ ██╔╝    ██╔══██╗██╔══██╗██║
██████╔╝██║   ██║██║   ██║█████╔╝     ███████║██████╔╝██║
██╔══██╗██║   ██║██║   ██║██╔═██╗     ██╔══██║██╔═══╝ ██║
██████╔╝╚██████╔╝╚██████╔╝██║  ██╗    ██║  ██║██║     ██║
╚═════╝  ╚═════╝  ╚═════╝ ╚═╝  ╚═╝    ╚═╝  ╚═╝╚═╝     ╚═╝`
	fmt.Println(banner)
	fmt.Println("┌───────────────────────────────────────────────────────┐")
	fmt.Println("│  Author: wimi                                         │")
	fmt.Println("│  Email: ywm00@qq.com                                  │")
	fmt.Println("│  Github: https://github.com/qmdx00                    │")
	fmt.Println("└───────────────────────────────────────────────────────┘")
}