# Sync同步模块(synchronization)
内存同步访问控制（原文是 memory access synchronization）
# 一. 互斥锁（同步)
特点

1. 原子性：把一个互斥量锁定为一个原子操作，这意味着操作系统（或pthread函数库）保证了如果一个线程锁定了一个互斥量，没有其他线程在同一时间可以成功锁定这个互斥量；

2. 唯一性：如果一个线程锁定了一个互斥量，在它解除锁定之前，没有其他线程可以锁定这个互斥量；

3. 非繁忙等待：如果一个线程已经锁定了一个互斥量，第二个线程又试图去锁定这个互斥量，则第二个线程将被挂起（不占用任何cpu资源），直到第一个线程解除对这个互斥量的锁定为止，第二个线程则被唤醒并继续执行，同时锁定这个互斥量。
# 二. 条件变量（同步）
	与互斥锁不同，条件变量是用来等待而不是用来上锁的。条件变量用来自动阻塞一个线程，直 到某特殊情况发生为止。通常条件变量和互斥锁同时使用。
	条件变量使我们可以睡眠等待某种条件出现。条件变量是利用线程间共享的全局变量进行同步 的一种机制，主要包括两个动作：
	一个线程等待"条件变量的条件成立"而挂起；
	另一个线程使 “条件成立”（给出条件成立信号）
【原理】：

	条件的检测是在互斥锁的保护下进行的。线程在改变条件状态之前必须首先锁住互斥量。如果一个条件为假，一个线程自动阻塞，并释放等待状态改变的互斥锁。
	如果另一个线程改变了条件，它发信号给关联的条件变量，唤醒一个或多个等待它的线程，重新获得互斥锁，重新评价条件。如果两进程共享可读写的内存，条件变量可以被用来实现这两进程间的线程同步
# 三. 读写锁（同步）

	读写锁与互斥量类似，不过读写锁允许更改的并行性，也叫共享互斥锁。互斥量要么是锁住状态，要么就是不加锁状态，而且一次只有一个线程可以对其加锁。读写锁可以有3种状态：读模式下加锁状态、写模式加锁状态、不加锁状态。

	一次只有一个线程可以占有写模式的读写锁，但是多个线程可以同时占有读模式的读写锁（允许多个线程读但只允许一个线程写）。

【读写锁的特点】：

	如果有其它线程读数据，则允许其它线程执行读操作，但不允许写操作；

	如果有其它线程写数据，则其它线程都不允许读、写操作。

【读写锁的规则】：

	如果某线程申请了读锁，其它线程可以再申请读锁，但不能申请写锁；

	如果某线程申请了写锁，其它线程不能申请读锁，也不能申请写锁。

	读写锁适合于对数据结构的读次数比写次数多得多的情况
# 四. 自旋锁（同步)

	自旋锁与互斥量功能一样，唯一一点不同的就是互斥量阻塞后休眠让出cpu，而自旋锁阻塞后不会让出cpu，会一直忙等待，直到得到锁。
	自旋锁在用户态使用的比较少，在内核使用的比较多！自旋锁的使用场景：锁的持有时间比较短，或者说小于2次上下文切换的时间。
	自旋锁在用户态的函数接口和互斥量一样，把pthread_mutex_xxx()中mutex换成spin，如：pthread_spin_init()
# 五. 信号量（同步与互斥）

	信号量广泛用于进程或线程间的同步和互斥，信号量本质上是一个非负的整数计数器，它被用来控制对公共资源的访问。

	编程时可根据操作信号量值的结果判断是否对公共资源具有访问的权限，当信号量值大于 0 时，则可以访问，否则将阻塞。
    PV 原语是对信号量的操作，一次 P 操作使信号量减１，一次 V 操作使信号量加１

