#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

void* do_print(void* unused) {
    sleep(10);
    printf("%s", "Hello, praktikum!");
    return 0;
}

int main(int argc, char** argv) {
    pthread_t tid;
    int status = pthread_create(&tid, 0, do_print, 0);
    if (status != 0) {
        printf("failed to create thread: got status %d\n", status);
        exit(1);
    }
    status = pthread_join(tid, 0);
    if (status != 0) {
        printf("failed to join thread: got status %d\n", status);
        exit(1);
    }
    return 0;
}
