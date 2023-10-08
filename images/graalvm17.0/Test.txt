class Test{
	public static void main(String[] args){
		System.out.println("This is only a Test!");
	}
}

//docker run -itd --name graalvm graalvm bash
//docker exec -it graalvm bash
//javac Test.java
//native-image Test //默认是.class
//native-image -jar springboot3 //编译的.jar文件是不完整的,没有aot