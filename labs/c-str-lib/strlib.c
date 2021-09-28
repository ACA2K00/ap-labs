

int mystrlen(char *str){
    int cont = 0;
    for (int i = 0; str[i] != '\0'; ++i){
        ++cont;
    }
    return cont;
}

char *mystradd(char *origin, char *addition){
    int cont = 0;
    static char res[100];
    for (int i = 0; origin[i] != '\0'; ++i){
        res[i]=origin[i];
        ++cont;
    }
    for (int i = 0; addition[i] != '\0'; ++i){
        res[cont]=addition[i];
        ++cont;
    }
    return res;
}

int mystrfind(char *origin, char *substr){
    int i, j, firstOcc;
   i = 0, j = 0;
 
   while (origin[i] != '\0') {
 
      while (origin[i] != substr[0] && origin[i] != '\0')
         i++;
 
      if (origin[i] == '\0')
         return (-1);
 
      firstOcc = i;
 
      while (origin[i] == substr[j] && origin[i] != '\0' && substr[j] != '\0') {
         i++;
         j++;
      }
 
      if (substr[j] == '\0')
         return (firstOcc);
      if (origin[i] == '\0')
         return (-1);
 
      i = firstOcc + 1;
      j = 0;
   }
}

int add(int a, int b) {
    return a + b;
}
