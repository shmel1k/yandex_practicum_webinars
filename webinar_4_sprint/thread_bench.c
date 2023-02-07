#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

void* do_print(void* unused) { return 0; }

#define MAX_THREADS_COUNT 100

int main(int argc, char** argv) {
    pthread_t* tids = (pthread_t*)calloc(MAX_THREADS_COUNT, sizeof(pthread_t));
    for (int i = 0; i < MAX_THREADS_COUNT; ++i) {
        int status = pthread_create(&tids[i], 0, do_print, 0); // Создали тред
        if (status != 0) {
            printf("failed to create thread: got status %d\n", status);
            exit(1);
        }
    }
    for (int i = 0; i < MAX_THREADS_COUNT; ++i) {
        int status = pthread_join(tids[i], 0); // Подождали завершения треда
        if (status != 0) {
            printf("failed to join thread: got status %d\n", status);
            exit(1);
        }
    }
    return 0;
}
