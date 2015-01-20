# Python 3.3.3 and 2.7.6
# python helloworld_python.py

from threading import Thread
from threading import Lock
i = 0

lock = Lock()


def someThreadFunction():
	lock.acquire()
	global i
	for j in range(1,1000000):
		i = i + 1
	lock.release()

def someAnotherThreadFunction():
	lock.acquire()
	global i
	for j in range(1,1000000):
		i = i - 1
	lock.release()	

def main():
	someThread = Thread(target = someThreadFunction, args = (),)
	someThread.start()
	
	someAnotherThread = Thread(target = someAnotherThreadFunction, args = (),)    
	someAnotherThread.start()
	
	someThread.join()
	someAnotherThread.join()
	
	print(i)
    
main()