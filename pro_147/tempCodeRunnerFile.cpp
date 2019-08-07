#include<bits/stdc++.h>
using namespace std;
int main()
{
	int n,m;
	cin>>n>>m;
	int a[n+1][m+1],b[n+1][m+1];
	for(int i=1;i<=n;i++)
	  for(int j=1;j<=m;j++)
	    {
	  	  cin>>a[i][j];
	  	  b[i][j]=a[i][j];
	    }  
	for(int i=2;i<=n-1;i++)
	  for(int j=2;j<=m-1;j++)
	  	b[i][j]=round((a[i][j]+a[i-1][j]+a[i+1][j]+a[i][j-1]+a[i][j+1])/5);
	for(int i=1;i<=n;i++)
	  {
	    for(int j=1;j<=m;j++)
		  cout<<b[i][j]<<" ";
		  cout<<endl;
	  }  

    // cout<<ceil(179/5);
}