// #include<stdio.h>
#include<bits/stdc++.h>
int main()
{

    int **a;
    int i, j,n;
    scanf("%d",&n);
    a = (int**)malloc(sizeof(int*)*n);//二维数组
    for (i = 0; i < n; ++i) //
    {
        a[i] = (int*)malloc(sizeof(int)*n);
    }
    for(i=0; i<n; i++)
     {
         a[i][0]=1;
         a[i][i]=1;

     }
     for(i=2; i<n; i++)
     {
         for(j=1; j<i; j++)
             a[i][j]=a[i-1][j-1]+a[i-1][j];
     }
     for(i=0; i<n; i++)
     {
         for(j=0; j<=i; j++)
        printf("%d ",a[i][j]);
        if(i!=n-1)
        printf("\n");

     }
return 0;
}
