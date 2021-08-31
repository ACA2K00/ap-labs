#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// adds/subtracts/multiplies all values that are in the *values array.
// nValues is the number of values you're reading from the array
// operator will indicate if it's an addition (1), subtraction (2) or
// multiplication (3)
long calc(int operator, int nValues, int *values) {
    long res;

    if(operator == 1){
        res = 0;
        for(int i = 0; i<nValues; i++){
            res += (long)values[i];
        }
    }else if(operator == 2){
        res = (long)values[0];
        for(int i = 1; i<nValues; i++){
            res -= (long)values[i];
        }
    }else{
        res = 1;
        for(int i = 0; i<nValues; i++){
            res *= (long)values[i];
        }
    }

    return res;
}

int main(int argc, char *argv[]) {
    int operator;
    int nValues = argc - 2;
    int values[nValues];
    int k = 0;
    char op[5];
    long result;

    strcpy(op, argv[1]);

    for(int i=2; i<argc; i++){
    int x = atoi(argv[i]);
    if(x == 0){
        printf("Invalid Number\n");
    }else{
        values[k] = x;
        //printf("x: %d\n",x);
        //printf("Value %d: %d\n", k + 1, values[k]);
        k++;
    }
    }

    if(strcmp(op, "add") == 0){
        operator = 1;
        result = calc(operator, nValues, values);
        printf("Result: %ld\n", result);
        //printf("Operator: %d\n", operator);
    }else if(strcmp(op, "sub") == 0){
        operator = 2;
        result = calc(operator, nValues, values);
        printf("Result: %ld\n", result);
        //printf("Operator: %d\n", operator);
    }else if(strcmp(op, "mult") == 0){
        operator = 3;
        result = calc(operator, nValues, values);
        printf("Result: %ld\n", result);
        //printf("Operator: %d\n", operator);
    }else{
        printf("Invalid Operator\n");
    }

    return -1;
}
