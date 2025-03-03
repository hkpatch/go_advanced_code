# AES高级加密标准（Advanced Encryption Standard）
又称Rijndael加密法，一个对称分组密码算法，是美国联邦政府采用的一种区块加密标准。这个标准用来替代原先的DES（Data Encryption Standard），已经被多方分析且广为全世界所使用。
AES中常见的有三种解决方案，分别为AES-128、AES-192和AES-256。

AES加密过程涉及到4种操作：字节替代（SubBytes）、行移位（ShiftRows）、列混淆（MixColumns）和轮密钥加（AddRoundKey）。
解密过程分别为对应的逆操作。由于每一步操作都是可逆的，按照相反的顺序进行解密即可恢复明文。
加解密中每轮的密钥分别由初始密钥扩展得到。算法中16字节的明文、密文和轮密钥都以一个4x4的矩阵表示。

如果采用真正的128位加密技术甚至256位加密技术，蛮力攻击要取得成功需要耗费相当长的时间。
1. 电码本模式（Electronic Codebook Book (ECB)）、

2. 密码分组链接模式（Cipher Block Chaining (CBC)）、

3. 计算器模式（Counter (CTR)）、

4. 密码反馈模式（Cipher FeedBack (CFB)）

5. 输出反馈模式（Output FeedBack (OFB)）

Note: 
- 出于安全考虑，golang默认并不支持ECB模式。
- AES对加密key的长度要求必须固定为16、24、32位，也就是128、192、256比特，所以又有一个AES-128、AES-192、AES-256这种叫法，位数越大安全性越高但加密速度越慢.
- 最关键是对明文长度也有要求，必须是分组长度长度的倍数，AES加密数据块分组长度必须为128bit也就是16位，所以这块又涉及到一个填充问题，而这个填充方式可以分为PKCS7和PKCS5等方式

## 背景
RSA公司举办过破译DES的比赛（DES Challenge)：

- 1997年的DES ChallengeI中用了96天；
- 1998年的DES ChallengeII-1中用了41天；
- 1998年的DES ChallengeII-2中用了56小时；
- 1999年的DES ChallengeII中用了22小时15分钟。


## 主要介绍：密码分组链接模式（Cipher Block Chaining (CBC)）

之所以叫这个名字，是因为密文分组像链条一样相互连接在一起。

在CBC模式中，每个明文块先与前一个密文块进行异或后，再进行加密。在这种方法中，每个密文块都依赖于它前面的所有明文块。
同时，为了保证每条消息的唯一性，在第一个块中需要使用初始化向量。
若第一个块的下标为1，则CBC模式的加密过程为： Ci = Ek (P ⊕ Ci-1), C0 = IV. 
而其解密过程则为： Pi = Dk (Ci) ⊕Ci-1, C0 = IV. 


## AES的Go实现
```go
// /usr/local/go/src/crypto/aes/cipher.go
func NewCipher(key []byte) (cipher.Block, error)
```
创建一个cipher.Block接口。参数key为密钥，长度只能是16、24、32字节，用以选择AES-128、AES-192、AES-256。

