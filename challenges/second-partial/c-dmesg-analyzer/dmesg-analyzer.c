#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include<fcntl.h>
#include<errno.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h> 
#include <unistd.h>

#define REPORT_FILE "report.txt"
#define HASHSIZE 10000

/*Alejandro Castro Arevalo-A01636235*/

/*HashTable-Start*/
/*Constructor of the nlist structure*/
struct nlist{
    struct nlist *next;
    char *name; /*text*/
    char *defn; /*Replacement text*/
};

static struct nlist *hashtab[HASHSIZE];

unsigned hash(char *s){
    unsigned hashval;
    for(hashval=0; *s!='\0'; s++) hashval = *s + 31 * hashval;
    return hashval%HASHSIZE; 
}

/*Methods of the nlist structure*/
struct nlist *lookup(char *s){
    struct nlist *np;
    for(np = hashtab[hash(s)]; np != NULL; np=np->next){
        if(strcmp(s, np->name) == 0){
            return np; /*found*/
        }
    }
    return NULL; /*not found*/
}

//struct nlist *lookup(char *);

struct nlist *install(char *name, char *defn){
    struct nlist *np;
    unsigned hashval;

    if((np = lookup(name)) == NULL){
        np = (struct nlist *) malloc(sizeof(*np));
        if(np == NULL || (np->name = strdup(name)) == NULL){
            return NULL;
        }
        hashval = hash(name);
        np->next = hashtab[hashval];
        hashtab[hashval] = np;
    }else{
        free((void *)np->defn);
    }
    if((np->defn = strdup(defn)) == NULL){
        return NULL;
    }
    return np;
}

/*Structure to hold the data of the Log*/
struct data{
    char *category;
    char *info;
};

/*data Method to sort the log*/
struct data arrangeData(char *s){
    struct data arranged;

    int str_length = 0;
    int cat_length = 0;
    int colons = 0;
    int blanks = 0;

    char *time; /*First part of the info (Before category)*/
    char *desc; /*Second part of the info (After Category)*/
    char *cat;
    char *cur = s; /*Current part of the logmsg*/

    while(*cur){
        str_length++;
        cur++;
        if(*(cur-1)==']') break; /*End of the timestamp*/
    }

    if(str_length<1){
        arranged.category = "General";
        arranged.info = (char*)malloc(strlen(s)*sizeof(char)+1);
        strcpy(arranged.info, s);
        return arranged;
    };

    time = (char*)malloc(str_length*sizeof(char));
    strncpy(time, s, str_length); /*Assign timestamp*/

    while(*cur){
        if(cat_length == 0 && *cur == ' '){
            blanks++;
        }else if(colons && *cur == ' '){
            break;
        }else if(!colons && *cur==':'){
            colons++;
            cat_length++;
        }else{
            cat_length++;
        }
        cur++;
    }

    if(colons){
        desc = (char*)malloc(strlen(s)*sizeof(char)-blanks-cat_length+1);
        strcpy(desc, time);
        strcat(desc, cur);

        cat = (char*)malloc(cat_length*sizeof(char));
        strncpy(cat, s+blanks+str_length, cat_length);
    }else{
        cat = (char*)malloc(strlen("General")*sizeof(char)+1);
        strcpy(cat,"General");
        desc = (char*)malloc(strlen(s)*sizeof(char)+1);
        strcpy(desc,s);
    }
    arranged.category = cat;
    arranged.info = desc;
    return arranged;
}
/*HashTable-End*/

void analizeLog(char *logFile, char *report);

int main(int argc, char **argv) {

    if (argc < 2) {
	printf("No text file recieved.\nCorrect Usage: ./dmesg-analizer logfile.txt\n");
	return 1;
    }

    analizeLog(argv[1], REPORT_FILE);

    return 0;
}

void analizeLog(char *logFile, char *report) {
    printf("Generating Report from: [%s] log file\n", logFile);

    int file = open(logFile, O_RDONLY);
	if (file == -1){
        printf("Error reading the log file\n");
        return;
    }

	int file_len = 0;
	char buff;
    int bytes_read;
    int newLine = 1;
	do {
        
        bytes_read = read(file, &buff, 1);
        if(newLine && bytes_read>0) {
            file_len++;
            newLine=0;
        }
        if(buff=='\n') newLine=1;
    } while (bytes_read>0);

    printf("Analyzing: %d lines\n",file_len);
    lseek(file,0,SEEK_SET);
    int line_length = 0;
    do{
        bytes_read = read(file, &buff, 1);
        if(bytes_read>0){
            if(buff =='\n'){
                if(line_length<1){
                    continue;
                }

                line_length++;
                lseek(file,-line_length,SEEK_CUR);
                char* lineBuff = (char*) malloc(line_length*sizeof(char));
                read(file,lineBuff,line_length);
                lineBuff[strcspn(lineBuff, "\n")] = 0;
                //printf("%s\n",lineBuff);
                line_length=0;
                struct data datos = arrangeData(lineBuff);
                install(datos.category,datos.info);
            }else{
                line_length++;
            }
        }
    } while(bytes_read>0);

    close(file);

    /*Write report*/
    FILE *result = fopen(report, "w+");
    struct nlist *head; /*Head of the hashtable*/

    for(int i=0; i<=HASHSIZE; i++){
        head=hashtab[i];
        if(head==NULL) continue;
        fputs(head->name, result);
        fputs("\n", result);
        struct nlist *cur;
        for(cur=head; cur!=NULL; cur=cur->next){
            fputs(" ", result);
            fputs(cur->defn, result);
            fputs("\n", result);
        }
    }
    fclose(result);

    printf("Report is generated at: [%s]\n", report);
}
