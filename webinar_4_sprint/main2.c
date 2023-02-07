#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

void* yield(void* unused) {
    printf("%s\n", "Hello, practicum!");
    fflush(stdout);
    sleep(10); // Подождём немножко.
    sched_yield(); // Запустим планировку процессов, оно же thread_yield.
    printf("%s\n", "I am back!"); // После выполнения остальных потоков наш поток будет опять обработан.
    return 0;
}

int main(int argc, char** argv) {
    pthread_t tid;
    int status = pthread_create(&tid, 0, yield, 0); // Создание нового потока.
    if (status != 0) {
        printf("failed to create thread: got status %d\n", status);
        exit(1);
    }
    status = pthread_join(tid, 0); // Ожидаем, пока поток завершится.
    /*
     * Важно. Данный поток будет заблокирован на выполнение,
     * даже когда за ним придут через thread_yield.
     */
    if (status != 0) {
        printf("failed to join thread: got status %d\n", status);
        exit(1);
    }
    printf("%s\n", "Exiting...");
    return 0;
}
