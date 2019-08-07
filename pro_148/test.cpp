#include <stdio.h>
#include <iostream>
using namespace std;
int a[100][100];
int main()
{
    int n;
    cin>>n;
    for(int i=0;i<=n;i++)
    for(int j=0;j<=n;j++)
    {
        a[i][j]=0;
    }
       for(int i=1;i<=n;i++)

    for(int j=1;j<=n;j++)
    {
        if(j==1&&j==1)
            a[i][j]=1;
        else
        a[i][j]=a[i-1][j-1]+a[i-1][j];
    }
    for(int i=1;i<=n;i++)
    {
        for(int j=1;j<=i;j++)
    {
        cout<<a[i][j];
        if(j!=i)
            cout<<' ';
    }
    if(i!=n)
    cout<<'\n';
    }
  
}

 return 0;   
}
  