#include <stdio.h>

int add(int, int);
int mystrlen(char *s);
char *mystradd(char *o, char *a);
int mystrfind(char *o, char *f);

int main() {
    printf("2+3=%d\n", add(2,3));
    printf("Blue = %d\n", mystrlen("Blue"));
    printf("Origin + Add = %s\n", mystradd("Origin", "Add"));
    printf("Azul Rojo Verde, Verde in: %d\n", mystrfind("Azul Rojo Verde", "Azul"));
    return 0;
}
