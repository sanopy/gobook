# ex 1.10

## Result

```bash
 $ go run fetchall.go https://time.is https://arxiv.org/pdf/1512.03385.pdf
2.17s   35786 https://time.is
7.32s  819383 https://arxiv.org/pdf/1512.03385.pdf
7.32s elapsed
 $ gobook/ch01/ex10 go run fetchall.go https://time.is https://arxiv.org/pdf/1512.03385.pdf
0.96s   35786 https://time.is
15.43s  819383 https://arxiv.org/pdf/1512.03385.pdf
15.43s elapsed
 $ gobook/ch01/ex10 diff arxiv.org_20210912212759.dump arxiv.org_20210912212815.dump
 $ gobook/ch01/ex10 diff time.is_20210912212800.dump time.is_20210912212815.dump
19c19
< _tD(1631449680738)
---
> _tD(1631449695799)
349c349
< </div><div id="time_section" class="w1"><div id="msgs" class="tr w90 vsbl"><div id="msgdiv"><h1 id="syncH" style="float:left">Exact time now:</h1></div><div id="syncDtl" class="w1">&nbsp;</div><div id="front_loc" class="w1" style="visibility:hidden"><a href="Kita,_Tōkyō,_Japan">北区, 東京都, 日本</a>での現在時刻:</div></div><div style="width:100%;position:relative"><div id="clock0_bg" onclick="clockclick(event)"><time id="clock">21:28:00</time></div></div><noscript><h2 class="w90 error">あなたのブラウザではJavaScriptが見つからないか、切られている ため、時計が動きません。</h2></noscript><div id="dd" class="w90 tr clockdate" onclick="location='/calendar'" title="クリックしてカレンダーを表示">2021年, 9月 12日, 日曜日, 第36週</div><div id="daydiv" class="w90 tr"><a href="https://www.un.org/en/observances/south-south-cooperation-day">United Nations Day for South-South Cooperation</a></div><div id="lC" class="w90 tr"><span id="locw">太陽 <span id="sun" class="nw">&uarr; 05:21 &darr; 17:54 (12時 32分)</span> <a href="Kita,_Tōkyō,_Japan#time_zone">詳細</a></span><div class="lsp w90 tr"></div></div><div class="w90 tr"><ul id="favs" class="tbx"><li><a href="/Tokyo" id="time-0">東京<br><span id="favt0">21:28</span></a></li><li><a href="/Beijing" id="time-1">北京市<br><span id="favt1">20:28</span></a></li><li><a href="/Moscow" id="time-2">モスクワ<br><span id="favt2">15:28</span></a></li><li><a href="/Paris" id="time-3">パリ<br><span id="favt3">14:28</span></a></li><li><a href="/London" id="time-4">ロンドン<br><span id="favt4">13:28</span></a></li><li><a href="/New_York" id="time-5">ニューヨーク<br><span id="favt5">08:28</span></a></li><li><a href="/Los_Angeles" id="time-6">ロサンゼルス<br><span id="favt6">05:28</span></a></li></ul></div><div id="recover" style="width:100%"><div class="tr" style="width:90%;padding:0 5% 20px 5%"><div class="rad" style="float:right;padding-top:0"><div id="time.is_728x90_970x90_970x250_300x250_320x50_ATF" style="position:relative;float:right"><script data-cfasync="false">
---
> </div><div id="time_section" class="w1"><div id="msgs" class="tr w90 vsbl"><div id="msgdiv"><h1 id="syncH" style="float:left">Exact time now:</h1></div><div id="syncDtl" class="w1">&nbsp;</div><div id="front_loc" class="w1" style="visibility:hidden"><a href="Kita,_Tōkyō,_Japan">北区, 東京都, 日本</a>での現在時刻:</div></div><div style="width:100%;position:relative"><div id="clock0_bg" onclick="clockclick(event)"><time id="clock">21:28:15</time></div></div><noscript><h2 class="w90 error">あなたのブラウザではJavaScriptが見つからないか、切られている ため、時計が動きません。</h2></noscript><div id="dd" class="w90 tr clockdate" onclick="location='/calendar'" title="クリックしてカレンダーを表示">2021年, 9月 12日, 日曜日, 第36週</div><div id="daydiv" class="w90 tr"><a href="https://www.un.org/en/observances/south-south-cooperation-day">United Nations Day for South-South Cooperation</a></div><div id="lC" class="w90 tr"><span id="locw">太陽 <span id="sun" class="nw">&uarr; 05:21 &darr; 17:54 (12時 32分)</span> <a href="Kita,_Tōkyō,_Japan#time_zone">詳細</a></span><div class="lsp w90 tr"></div></div><div class="w90 tr"><ul id="favs" class="tbx"><li><a href="/Tokyo" id="time-0">東京<br><span id="favt0">21:28</span></a></li><li><a href="/Beijing" id="time-1">北京市<br><span id="favt1">20:28</span></a></li><li><a href="/Moscow" id="time-2">モスクワ<br><span id="favt2">15:28</span></a></li><li><a href="/Paris" id="time-3">パリ<br><span id="favt3">14:28</span></a></li><li><a href="/London" id="time-4">ロンドン<br><span id="favt4">13:28</span></a></li><li><a href="/New_York" id="time-5">ニューヨーク<br><span id="favt5">08:28</span></a></li><li><a href="/Los_Angeles" id="time-6">ロサンゼルス<br><span id="favt6">05:28</span></a></li></ul></div><div id="recover" style="width:100%"><div class="tr" style="width:90%;padding:0 5% 20px 5%"><div class="rad" style="float:right;padding-top:0"><div id="time.is_728x90_970x90_970x250_300x250_320x50_ATF" style="position:relative;float:right"><script data-cfasync="false">
```
