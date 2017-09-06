/*Task

Write a Person class with an instance variable, , and a constructor that takes
an integer, , as a parameter.
The constructor must assign  to  after confirming the argument passed as is not
negative; if a negative argument is
passed as , the constructor should set  to  and print Age is not valid, setting
age to 0.. In addition, you must write
the following instance methods:

yearPasses() should increase the  instance variable by .

amIOld() should perform the following conditional actions:

If , print You are young..
If  and , print You are a teenager..
Otherwise, print You are old..

 */
package person;

import java.io.*;
import java.util.*;

public class Person {
    private int age;

    public Person(int initialAge) {

        if(initialAge < 0){
            this.age = 0;
            System.out.println("Age is not valid, setting age to 0.");
        }

        this.age = initialAge;

}

    public void amIOld() {

        if(this.age < 13){
            System.out.println("You are young.");
        }

        else if(this.age >=13 && this.age < 18){
            System.out.println("You are a teenager.");
        }

        else{
            System.out.println("You are old.");
        }

}

    public void yearPasses() {
        this.age ++;
}

    public static void main(String[] args) {

    }

}
