package cmd

import (
	"runtime"

	"github.com/urfave/cli/v2"
)

var JCT = newJCT()

// newJCT 构造工具集
func newJCT() *cli.App {
	app := cli.NewApp()
	app.Usage = "命令行工具集"
	app.Commands = append(app.Commands, newCRLF())
	return app
}

// newCRLF 构造crlf命令
func newCRLF() *cli.Command {
	c := &cli.Command{}
	c.Name = "crlf"
	c.Usage = "转换文本换行符"
	c.Action = docrlf
	c.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "type",
			Aliases:  []string{"t"},
			Usage:    "换行符类型(cr,lf,crlf)",
			Required: true,
			Value: func() string {
				switch runtime.GOOS {
				case "windows":
					return "crlf"
				default:
					return "lf"
				}
			}(),
		},
		&cli.BoolFlag{
			Name:    "dir",
			Aliases: []string{"d"},
			Usage:   "使用文件夹",
		},
	}
	return c
}

// docrlf 处理换行符
func docrlf(c *cli.Context) error {
	fs := c.Args().Slice()
	t := c.String("t")
	if c.Bool("d") {
		// 处理文件夹
		for _, dir := range fs {
			err := procDir(t, dir)
			if err != nil {
				return err
			}
		}
	} else {
		// 处理文件
		for _, file := range fs {
			err := procFile(t, file)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
