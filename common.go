package main

import (
	"fmt"
	"github.com/gookit/color"
)

const banner = `
  _________                        ___________              .__  .__       .__     
 /   _____/__ ________   __________\_   _____/ ____    ____ |  | |__| _____|  |__  
 \_____  \|  |  \____ \_/ __ \_  __ \    __)_ /    \  / ___\|  | |  |/  ___/  |  \ 
 /        \  |  /  |_> >  ___/|  | \/        \   |  \/ /_/  >  |_|  |\___ \|   Y  \
/_______  /____/|   __/ \___  >__| /_______  /___|  /\___  /|____/__/____  >___|  /
        \/      |__|        \/             \/     \//_____/              \/     \/ 

本程序用于英语词汇训练，支持手动添加词汇，并可以从添加的词汇表中随机抽查100个，支持随机中译英以及英译中。
				苟有恒，何必三更起五更眠；
				最无益，莫过一日曝十日寒。
`
const Author = `
Author:a1phaboy
`

const options = `
1. 添加词汇
2. 测试
3. 退出
`
const bash =`SuperEnglish > `

func showBanner(){
	color.RGBStyleFromString("0,255,0").Println(banner)
	color.RGBStyleFromString("0,255,0").Println(Author)
	color.RGBStyleFromString("0,255,0").Println(options)
}
func showBash(){
	fmt.Print(bash)
}

func checkErr(err error){
	if err != nil{
		panic(err)
	}
}
