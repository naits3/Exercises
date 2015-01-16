# Python 3.3.3 and 2.7.6
# python helloworld_python.py

from threading import Thread

i = 0

def someThreadFunction():
    for i in range(1,1000000):
    	i = i + 1

def someAnotherThreadFunction():
	for i in range(1,1000000):
		i = i - 1
		
def main():
	global i
	someThread = Thread(target = someThreadFunction, args = (),)
	someAnotherThread = Thread(target = someAnotherThreadFunction, args = (),)    
	
	someThread.start()
	someThread.join()
	
	someAnotherThread.start()
	someAnotherThread.join()
    
	print("Hello from main!")
	print(i)
    
main()
