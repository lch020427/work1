package main

import "fmt"
var M,N,T,Sy,Sx,Ex,Ey,x,y int
var area [][]int

func main(){
	fmt.Scanf("%d%d%d",&M,&N,&T)
	fmt.Scanf("%d%d%d%d",&Sx,&Sy,&Ex,&Ey)
	for i:=0;i<M;i++{
		for j:=0;j<N;i++{
			area[i][j]=1
		}
	}
	for i:=0;i<M;i++{
		for j:=0;j<N;j++{
			fmt.Scanln(&x,&y)
			area[x][y]=0
		}
	}
	Go(Sx,Sy,Ex,Ey)
}
func Go(Sx,Sy,Ex,Ey int){
	var flag int
	if(Sy==Ey&&Sx==Ex){
		flag=1
	}
	if(flag!=1&&area[Sx-1][Sy]==1){
		Go(Sx-1,Sy,Ex,Ey)
		fmt.Printf("%d %d",Sx-1,Sy)
	}else if(flag!=1&&area[Sx+1][Sy]==1){
		Go(Sx+1,Sy,Ex,Ey)
		fmt.Println(Sx+1,Sy)
	}else if(flag!=1&&area[Sx][Sy-1]==1){
		Go(Sx,Sy-1,Ex,Ey)
		fmt.Println(Sx,Sy-1)
	}else if(flag!=1&&area[Sx][Sy+1]==1){
		Go(Sx,Sy+1,Ex,Ey)
		fmt.Println(Sx,Sy+1)
	}
}