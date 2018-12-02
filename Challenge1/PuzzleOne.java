// used this tutorial for reading file line by line https://www.roseindia.net/java/beginners/java-read-file-line-by-line.shtml
import java.io.*;

public class PuzzleOne {
  public int frequency = 0;
  public String fileName;

  public PuzzleOne(String fileName) {
    this.setFileName(fileName);
  }

  private void setFileName(String name) {
    this.fileName = name;
  }

  private String getFileName() {
    return this.fileName;
  }

  private void setFrequency(int change, char operation) {
    if(operation == '+') {
      this.frequency += change;
    } else {
      this.frequency -= change;
    }
  }

  private int getFrequency() {
    return this.frequency;
  }

  public int findFrequency() {
    try { 
      FileInputStream fstream = new FileInputStream(getFileName());
      DataInputStream in = new DataInputStream(fstream);
      BufferedReader reader = new BufferedReader(new InputStreamReader(in));
      String line;
      int frequency = 0;

      while((line = reader.readLine()) != null) {
        int frequencyChange = Integer.parseInt(line.substring(1, line.length()));
        setFrequency(frequencyChange, line.charAt(0)); 
      }
      in.close();
    } catch (FileNotFoundException e) {
      System.err.println("Error: " + e.getMessage());
    } catch (IOException e) {
      System.err.println("Error: " + e.getMessage());
    }
    return getFrequency();
  }


  public static void main(String[] args) {
    PuzzleOne puz = new PuzzleOne("1_input.txt");
    int currentFrequency = puz.findFrequency();
    System.out.println("CurrentFrequency: " + currentFrequency);
  }
}
