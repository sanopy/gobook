# ex 1.11

## Generate args

```js
> Array.from(document.querySelectorAll('.td.DescriptionCell a'), a => 'https://' + a.innerHTML).join(' ');
'https://Google.com https://Youtube.com https://Tmall.com https://Baidu.com https://Qq.com https://Sohu.com https://Facebook.com https://Taobao.com https://360.cn https://Jd.com https://Amazon.com https://Yahoo.com https://Wikipedia.org https://Weibo.com https://Sina.com.cn https://Zoom.us https://Xinhuanet.com https://Live.com https://Reddit.com https://Netflix.com https://Instagram.com https://Microsoft.com https://Office.com https://Google.com.hk https://Panda.tv https://Zhanqi.tv https://Alipay.com https://Bing.com https://Csdn.net https://Myshopify.com https://Vk.com https://Yahoo.co.jp https://Bongacams.com https://Microsoftonline.com https://Naver.com https://Twitch.tv https://Twitter.com https://Okezone.com https://Ebay.com https://Aparat.com https://Adobe.com https://Amazon.in https://Aliexpress.com https://Yy.com https://Huanqiu.com https://Tianya.cn https://Chaturbate.com https://Amazon.co.jp https://Linkedin.com https://Stackoverflow.com'
```

## Result

```bash
$ go run fetchall.go https://Google.com https://Youtube.com https://Tmall.com https://Baidu.com https://Qq.com https://Sohu.com https://Facebook.com https://Taobao.com https://360.cn https://Jd.com https://Amazon.com https://Yahoo.com https://Wikipedia.org https://Weibo.com https://Sina.com.cn https://Zoom.us https://Xinhuanet.com https://Live.com https://Reddit.com https://Netflix.com https://Instagram.com https://Microsoft.com https://Office.com https://Google.com.hk https://Panda.tv https://Zhanqi.tv https://Alipay.com https://Bing.com https://Csdn.net https://Myshopify.com https://Vk.com https://Yahoo.co.jp https://Bongacams.com https://Microsoftonline.com https://Naver.com https://Twitch.tv https://Twitter.com https://Okezone.com https://Ebay.com https://Aparat.com https://Adobe.com https://Amazon.in https://Aliexpress.com https://Yy.com https://Huanqiu.com https://Tianya.cn https://Chaturbate.com https://Amazon.co.jp https://Linkedin.com https://Stackoverflow.com
0.59s  207724 https://Stackoverflow.com
1.01s  239746 https://Chaturbate.com
Get "https://Microsoftonline.com": dial tcp: lookup Microsoftonline.com on 172.27.0.1:53: no such host
1.13s    9078 https://Myshopify.com
1.26s     173 https://Baidu.com
1.32s   76387 https://Twitter.com
1.45s  158040 https://Panda.tv
Get "https://Huanqiu.com": remote error: tls: handshake failure
1.75s  100257 https://Zoom.us
1.80s  111053 https://Twitch.tv
1.95s   37324 https://Yahoo.co.jp
1.95s       0 https://Instagram.com
1.96s   76406 https://Bing.com
1.98s   15148 https://Google.com
2.13s       0 https://Csdn.net
2.13s   36928 https://Live.com
2.26s  410511 https://Amazon.co.jp
2.28s  224422 https://Facebook.com
2.28s  217627 https://Sohu.com
2.67s  543622 https://Youtube.com
2.67s  389277 https://Bongacams.com
3.65s  137101 https://Tmall.com
4.25s  122075 https://Qq.com
4.41s    7624 https://Tianya.cn
4.59s  272667 https://Yy.com
4.70s  114038 https://Office.com
4.90s  531965 https://Sina.com.cn
5.00s  105809 https://Linkedin.com
5.01s   15553 https://Google.com.hk
5.01s   73675 https://Wikipedia.org
5.15s  155355 https://Ebay.com
5.82s  501231 https://Amazon.in
5.85s  139918 https://Adobe.com
5.86s  224310 https://Microsoft.com
5.96s  215759 https://Naver.com
5.99s   40030 https://Aliexpress.com
6.24s  402077 https://Amazon.com
6.41s  152839 https://Zhanqi.tv
6.63s  415699 https://Netflix.com
6.63s  147973 https://Aparat.com
6.66s  111195 https://Okezone.com
6.70s   20592 https://Jd.com
6.92s   54830 https://Vk.com
7.09s  895243 https://Reddit.com
7.26s   99044 https://Weibo.com
7.93s  582197 https://Yahoo.com
9.63s    9040 https://Taobao.com
18.50s   79789 https://360.cn
Get "https://Alipay.com": dial tcp 110.75.129.5:443: i/o timeout
Get "https://Xinhuanet.com": dial tcp 202.108.119.193:443: i/o timeout
30.00s elapsed
```
