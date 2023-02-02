# go-DripLoader

DripLoader 的 golang 实现

### 使用方法

msf 和 cs 直接生成的 shellcode 太小，我用原版的DripLoader也没成功，这里可以先生成一个dll，然后使用 sRDI 项目转换成 shellcode （我自己写了一个golang版本的，也可以用）：

```
// https://github.com/monoxgas/sRDI
python CovertToShellcode.py evil.dll
```

转换成功后，会生成一个 bin 文件，编译 go-DripLoader 后运行如下命令：

```
go-DripLoader.exe [pid] [binfile]
```


![ed72b1dc41f8a663a9b4a137bfb1a5d](https://user-images.githubusercontent.com/84751437/216241687-05ac43d1-dbd6-40f4-a3df-86b601193b74.png)



### 参考

 [Bypassing EDR Real-Time Injection Detection Logic - RedBluePurple](https://blog.redbluepurple.io/windows-security-research/bypassing-injection-detection) 

 [xuanxuan0/DripLoader: Evasive shellcode loader for bypassing event-based injection detection (PoC) (github.com)](https://github.com/xuanxuan0/DripLoader) 
