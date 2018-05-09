---


---

<h1 id="poetry-desktop">poetry-desktop</h1>
<p>欢迎使用诗词桌面！</p>
<p>程序使用 Go 语言开发。</p>
<p>目前仅支持 Windows 。</p>
<h1 id="简介">简介</h1>
<p>能够自动、定时地从网路上随机获取一首唐、宋诗或宋词，并将文本整合在背景图片上，最后自动设置为壁纸。</p>
<p>要查看帮助请在程序所在目录打开cmd命令行输入: poetry-desktop --help</p>
<pre><code>NAME:                                                                                           
   poetry-desktop.exe - A new cli application                                                   
                                                                                                
USAGE:                                                                                          
   poetry-desktop.exe [global options] command [command options] [arguments...]                 
                                                                                                
VERSION:                                                                                        
   1.0                                                                                        
                                                                                                
COMMANDS:                                                                                       
     help, h  Shows a list of commands or help for one command                                  
                                                                                                
GLOBAL OPTIONS:                                                                                 
   --interval value, -i value  设置诗词桌面自动切换时间间隔，单位：小时 (default: "1")                              
   --font value, -f value      设置显示的字体 (default: "fonts/FangZhengKaiTiJianTi-1.ttf")            
   --fontsize value, -s value  设置字体大小 (default: "25")                                           
   --imgloc value, -l value    指定背景图片位置 (default: "images/poetry_bg_0.jpg")                     
   --color value, -c value     设置字体颜色，格式为 rgba (default: "(0,0,0,0.8)")                         
   --xoffset value             设置文字整体在X轴方向的偏移量 (default: "+0")                                  
   --yoffset value             设置文字整体在Y轴方向的偏移量 (default: "+0")                                  
   --help, -h                  show help                                                        
   --version, -v               print the version                                
</code></pre>
<h1 id="使用效果">使用效果</h1>
<p><img src="http://128.199.227.220:8003/Temp/Screenshot%20from%202018-05-06%2021-26-24.png" alt="alt text"><br>
<img src="http://128.199.227.220:8003/Temp/Screenshot%20from%202018-05-06%2021-34-58.png" alt="alt text"></p>
<h1 id="下载">下载</h1>
<ol>
<li>请前往 <a href="https://github.com/okcy1016/poetry-desktop/releases">Releases</a> 页面下载（国内环境下速度可能较慢）</li>
<li><a href="https://yadi.sk/d/MhkJwXQ43VUaug">Yandisk</a></li>
</ol>
<h1 id="致谢">致谢</h1>
<p>感谢以下项目的贡献者：</p>
<ol>
<li>
<p>中华古诗词数据库<br>
<a href="https://github.com/chinese-poetry/chinese-poetry">https://github.com/chinese-poetry/chinese-poetry</a></p>
</li>
<li>
<p>A simple, fast, and fun package for building command line apps in Go<br>
<a href="https://github.com/urfave/cli">https://github.com/urfave/cli</a></p>
</li>
<li>
<p>ImageMagick 7<br>
<a href="https://github.com/ImageMagick/ImageMagick">https://github.com/ImageMagick/ImageMagick</a></p>
</li>
<li>
<p>Set the desktop wallpaper on Windows<br>
<a href="https://github.com/sindresorhus/win-wallpaper">https://github.com/sindresorhus/win-wallpaper</a></p>
</li>
</ol>

