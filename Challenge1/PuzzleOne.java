// used this tutorial for reading file line by line https://www.roseindia.net/java/beginners/java-read-file-line-by-line.shtml
import java.io.*;

public class PuzzleOne {
  public static void main(String[] args) {
    try { 
      FileInputStream fstream = new FileInputStream("1_input.txt");
      DataInputStream in = new DataInputStream(fstream);
      BufferedReader reader = new BufferedReader(new InputStreamReader(in));
      String line;
      int frequency = 0;

      while((line = reader.readLine()) != null) {
        int frequencyChange = Integer.parseInt(line.substring(1, line.length()));
        if(line.charAt(0) == '+') {
          frequency += frequencyChange;
        } else {
          frequency -= frequencyChange;
        }
      }
      in.close();
      System.out.println("Frequency: " + frequency);

    } catch (FileNotFoundException e) {
      System.err.println("Error: " + e.getMessage());
    } catch (IOException e) {
      System.err.println("Error: " + e.getMessage());
    }
  }
}
