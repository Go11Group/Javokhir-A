#include <stdio.h>
#include <time.h>
#include <stdlib.h>
void Random_Array(int nums[],int n){
	for(int i = 0; i<n ;i++){
           nums[i]= 14 + rand()%22;
		printf("%d ",nums[i]);
		}
	puts("");
	}

void Print_Array(int nums[],int n){
 	for(int i = 0 ; i<n ; i++){
	if(nums[i]%2==0){
	    printf(" + ");	
	}
	else printf("%d ",nums[i]);	
	
		}
	
}



int main(){
int n;

	srand(time(0));
	
	printf("n = ");scanf("%d",&n);
	int array[n];

   	Random_Array(array,n);
	Print_Array(array,n);

return 0;
}
