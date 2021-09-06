#include <stdio.h>
#include <stdlib.h>
#include <string.h>

//Alejandro Castro Arévalo-A01636235
/* month_day function's prototype*/
void month_day(int year, int yearday, int *pmonth, int *pday){
    int last_days[13] = {0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
    int leap_days[13] = {0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31};
    int days = last_days[0];
    int month;
    int day;
    static char *name[] = {
       "Illegal month",
       "January", "February", "March",
       "April", "May", "June",
       "July", "August", "September",
       "October", "November", "December"
    };

    if(((year%4==0) && ((year%400==0) || (year%100!=0)))){
        if(1 <= yearday && yearday <= 31){
            month = *(pmonth + 1);
            for(int i=0; i < month; i++){
                days += leap_days[i];
            }
            day = *(pday + (yearday-days));
        }else if (32 <= yearday && yearday <= 60){
            month = *(pmonth + 2);
            for(int i=0; i < month; i++){
                days += leap_days[i];
            }
            day = *(pday + (yearday-days));
        }else if (61 <= yearday && yearday <= 91){
            month = *(pmonth + 3);
            for(int i=0; i < month; i++){
                days += leap_days[i];
            }
            day = *(pday + (yearday-days));
        }else if (92 <= yearday && yearday <= 121){
            month = *(pmonth + 4);
            for(int i=0; i < month; i++){
                days += leap_days[i];
            }
            day = *(pday + (yearday-days));
        }else if (122 <= yearday && yearday <= 152){
            month = *(pmonth + 5);
            for(int i=0; i < month; i++){
                days += leap_days[i];
            }
            day = *(pday + (yearday-days));
        }else if (153 <= yearday && yearday <= 182){
            month = *(pmonth + 6);
            for(int i=0; i < month; i++){
                days += leap_days[i];
            }
            day = *(pday + (yearday-days));
        }else if (183 <= yearday && yearday <= 213){
            month = *(pmonth + 7);
            for(int i=0; i < month; i++){
                days += leap_days[i];
            }
            day = *(pday + (yearday-days));
        }else if (214 <= yearday && yearday <= 244){
            month = *(pmonth + 8);
            for(int i=0; i < month; i++){
                days += leap_days[i];
            }
            day = *(pday + (yearday-days));
        }else if (245 <= yearday && yearday <= 274){
            month = *(pmonth + 9);
            for(int i=0; i < month; i++){
                days += leap_days[i];
            }
            day = *(pday + (yearday-days));
        }else if (275 <= yearday && yearday <= 305){
            month = *(pmonth + 10);
            for(int i=0; i < month; i++){
                days += leap_days[i];
            }
            day = *(pday + (yearday-days));
        }else if (306 <= yearday && yearday <= 335){
            month = *(pmonth + 11);
            for(int i=0; i < month; i++){
                days += leap_days[i];
            }
            day = *(pday + (yearday-days));
        }else if (336 <= yearday && yearday <= 366){
            month = *(pmonth + 12);
            for(int i=0; i < month; i++){
                days += leap_days[i];
            }
            day = *(pday + (yearday-days));
        }else{
            month = *(pmonth + 0);
        }
    }else{
        if(1 <= yearday && yearday <= 31){
            month = *(pmonth + 1);
            for(int i=0; i < month; i++){
                days += last_days[i];
            }
            day = *(pday + (yearday-days));
        }else if (32 <= yearday && yearday <= 59){
            month = *(pmonth + 2);
            for(int i=0; i < month; i++){
                days += last_days[i];
            }
            day = *(pday + (yearday-days));
        }else if (60 <= yearday && yearday <= 90){
            month = *(pmonth + 3);
            for(int i=0; i < month; i++){
                days += last_days[i];
            }
            day = *(pday + (yearday-days));
        }else if (91 <= yearday && yearday <= 120){
            month = *(pmonth + 4);
            for(int i=0; i < month; i++){
                days += last_days[i];
            }
            day = *(pday + (yearday-days));
        }else if (121 <= yearday && yearday <= 151){
            month = *(pmonth + 5);
            for(int i=0; i < month; i++){
                days += last_days[i];
            }
            day = *(pday + (yearday-days));
        }else if (152 <= yearday && yearday <= 181){
            month = *(pmonth + 6);
            for(int i=0; i < month; i++){
                days += last_days[i];
            }
            day = *(pday + (yearday-days));
        }else if (182 <= yearday && yearday <= 212){
            month = *(pmonth + 7);
            for(int i=0; i < month; i++){
                days += last_days[i];
            }
            day = *(pday + (yearday-days));
        }else if (213 <= yearday && yearday <= 243){
            month = *(pmonth + 8);
            for(int i=0; i < month; i++){
                days += last_days[i];
            }
            day = *(pday + (yearday-days));
        }else if (244 <= yearday && yearday <= 273){
            month =*(pmonth + 9);
            for(int i=0; i < month; i++){
                days += last_days[i];
            }
            day = *(pday + (yearday-days));
        }else if (274 <= yearday && yearday <= 304){
            month = *(pmonth + 10);
            for(int i=0; i < month; i++){
                days += last_days[i];
            }
            day = *(pday + (yearday-days));
        }else if (305 <= yearday && yearday <= 334){
            month = *(pmonth + 11);
            for(int i=0; i < month; i++){
                days += last_days[i];
            }
            day = *(pday + (yearday-days));
        }else if (335 <= yearday && yearday <= 365){
            month = *(pmonth + 12);
            for(int i=0; i < month; i++){
                days += last_days[i];
            }
            day = *(pday + (yearday-days));
        }else{
            month = *(pmonth + 0);
        }
    }

    if(month == 0){
        printf("%s, Invalid Day input\n", name[month]);
    }else{
        if(yearday == 366){
            printf("%s %d, %d\n", name[month], 31, year);
        }else{
            printf("%s %d, %d\n", name[month], day-1, year);
        }
    }
}

int main(int argc, char *argv[]) {
    int year = atoi(argv[1]);
    int yearday = atoi(argv[2]);
    int month_number[13] = {0,1,2,3,4,5,6,7,8,9,10,11,12};
    int day_number[31];

    for(int i=0; i<31; i++){
        day_number[i] = i+1;
        //printf("Pos: %d, Val: %d \n", i, day_number[i]);
    }

    month_day(year, yearday, &month_number[0], &day_number[0]);
}
