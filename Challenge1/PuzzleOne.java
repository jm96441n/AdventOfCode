// used this tutorial for reading file line by line https://www.roseindia.net/java/beginners/java-read-file-line-by-line.shtml
import java.io.*;
import java.util.*;

public class PuzzleOne {
  private int frequency = 0;
  private int dupeFrequency;
  private String fileName;
  private HashMap<Integer, Integer> frequencies = new HashMap<Integer, Integer>();
  private ArrayList<String> readList = new ArrayList<String>();

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

  private HashMap<Integer, Integer> getFrequencies() {
    return this.frequencies;
  }

  public boolean hasNotSeenFrequency() {
    return !frequencies.containsKey(getFrequency());
  }

  public boolean hasSeenFrequency() {
    return frequencies.containsKey(getFrequency());
  }

  public void parseLine(String line) {
    int frequencyChange = Integer.parseInt(line.substring(1, line.length()));
    setFrequency(frequencyChange, line.charAt(0));
  }

  public int findDuplicateFrequency() {
    while(dupeFrequency == 0) {
      findFrequency();
    }
    return dupeFrequency;
  }

  public int findFrequency() {
    if(readList.isEmpty()) { parseFile(); }
    for(String line : readList) {
      parseLine(line);
      if(hasSeenFrequency()) {
         dupeFrequency = getFrequency();
         break;
       } else if(hasNotSeenFrequency()) {
         frequencies.put(getFrequency(), 1);
       }
     }
     return getFrequency();
  }

  public void parseFile() {
    try {
      FileInputStream fstream = new FileInputStream(getFileName());
      BufferedReader reader = new BufferedReader(new InputStreamReader(fstream));
      String line;
      while((line = reader.readLine()) != null) {
        readList.add(line);
      }
    } catch (FileNotFoundException e) {
      System.err.println("Error: " + e.getMessage());
    } catch (IOException e) {
      System.err.println("Error: " + e.getMessage());
    }
  }

  public static void main(String[] args) {
    PuzzleOne puz = new PuzzleOne("1_input.txt");
    int frequency = puz.findDuplicateFrequency();
    System.out.println("DuplicateFrequency: " + frequency);
  }
}
