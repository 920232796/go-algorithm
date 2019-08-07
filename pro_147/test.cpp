#include<iostream>
using namespace std;
int main() {
	int a[6][6],t,t1=0,t2=0,x,y;
	for(int i=1; i<=5; i++)
		for(int j=1; j<=5; j++)
			cin>>a[i][j];
	for(int i=1; i<=5; i++) {
        cout<<"i = "<<i<<endl;
		// for(int j=1; j<=5; j++) {
        //     cout<<"i = "<<i<<endl;
		// 	if(t<a[i][j]) {
		// 		t=a[i][j];
		// 		x=i;
		// 		y=j;
		// 	}
        // }
		for(int k=1; k<=5; k++)
			if(t>a[k][y])
				t1=1;
        cout<<"t1 = " <<t1<<endl;
		if(t1==0) {
			cout<<x<<" "<<y<<" "<<t;
			t2=1;
		}
		t=0;
		t1=0;
	}
	if(t2==0)
		cout<<"not found";
}