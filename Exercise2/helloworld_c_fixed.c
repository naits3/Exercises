// gcc 4.7.2 +
// gcc -std=gnu99 -Wall -g -o helloworld_c_fixed helloworld_c_fixed.c -lpthread

#include <pthread.h>
#include <stdio.h>


int i = 0;
pthread_mutex_t lock;


// Note the return type: void*
void* thread_1_function(){

	for(int n = 0;n<1000000;n++){
		
    	pthread_mutex_lock(&lock);
		i++;
		//printf("%i\n",i);
		pthread_mutex_unlock(&lock);
	}
	
    return NULL;
}

void* thread_2_function(){
    
	for(int n = 0;n<1000002;n++){
		pthread_mutex_lock(&lock);
		i--;
		//printf("%i\n",i);
		pthread_mutex_unlock(&lock);
	}
    return NULL;
}

int main(){

    pthread_t thread_1;
    pthread_create(&thread_1, NULL, thread_1_function, NULL);
    
	pthread_t thread_2;
    pthread_create(&thread_2, NULL, thread_2_function, NULL);
    
    pthread_join(thread_1, NULL);
	pthread_join(thread_2, NULL);

	pthread_mutex_destroy(&lock);
    
	printf("%i\n",i);
    return 0;
}
