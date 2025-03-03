# 备忘录模式【Memento】

## 意图：
	在不破坏封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。这样以后就可将该对象恢复到原先保存的状态。
	在备忘录模式中，如果要保存的状态多，可以创造一个备忘录管理者角色来管理备忘录。
	
## 适用性：
	必须保存一个对象在某一个时刻的(部分)状态, 这样以后需要时它才能恢复到先前的状态。
	如果一个用接口来让其它对象直接得到这些状态，将会暴露对象的实现细节并破坏对象的封装性

## 比如：

	1、需要保存和恢复数据的相关状态场景。如保存游戏状态的场景：撤销场景，事务回滚等；
	2、副本监控场景。备忘录可以当做一个临时的副本监控，实现非实时和准实时的监控。

## 需求：
	某线上博客平台, 需为用户提供在线编辑文章功能，文章主要包括标题 - title 和内容 - content等信息，
	为最大程度防止异常情况导致编辑内容的丢失, 需要提供版本暂存和Undo, Redo功能。
	"版本暂存"问题可以应用备忘录模式, 将编辑器的状态完整保存(主要就是编辑内容)，Undo和Redo的本质, 是在历史版本中前后移动

## 实现：
	IEditor: 定义编辑器接口
	tEditorMemento: 定义编辑器的备忘录, 也就是编辑器的内部状态数据模型, 同时也对应一个历史版本
	tMockEditor: 虚拟的编辑器类, 实现IEditor接口

## 角色
    发起人: 发起人的内部要规定要备忘的范围，负责提供备案数据
    备忘录: 存储发起人对象的内部状态，在需要的时候，可以向其他人提供这个内部状态，以方便负责人恢复发起人状态
    负责人: 负责对备忘录进行管理（保存或提供备忘录）

## 优点：
	1）简化发起人实体类（Originator）的职责，隔离状态存储与获取，
    实现了信息的封装，客户端无须关心状态的保存细节。
	2）提供状态回滚功能。

## 缺点：
	备忘录模式的缺点主要是消耗资源。
	如果需要保存的状态过多，则每一次保存都会消耗很多内存。
