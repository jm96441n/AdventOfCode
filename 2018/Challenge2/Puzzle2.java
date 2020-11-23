import java.io.*;
import java.util.*;

public class Puzzle2 {
  private ArrayList<String> lines = new ArrayList<String>();
  private int doublesCount = 0;
  private int triplesCount = 0;

  public void parseFile() {
    try {
      FileInputStream fstream = new FileInputStream("input.txt");
      BufferedReader reader = new BufferedReader(new InputStreamReader(fstream));
      String line;
      while((line = reader.readLine()) != null) {
        lines.add(line);
      }
    } catch (FileNotFoundException e) {
      System.err.println("Error: " + e.getMessage());
    } catch (IOException e) {
      System.err.println("Error: " + e.getMessage());
    }
  }

  public void findDoublesAndTriples() {
    for(String ele : lines) {
      HashMap<Character, Integer> charCounter = new HashMap<Character, Integer>();
      for(int i = 0; i < ele.length(); i++) {
        char currentChar = ele.charAt(i);
        if(charCounter.containsKey(currentChar)) {
          Integer currentVal = charCounter.get(currentChar);
          charCounter.replace(currentChar, ++currentVal);
        } else {
          charCounter.put(currentChar, 1);
        }
      }
      if(charCounter.containsValue(2)) { doublesCount++; } 
      if(charCounter.containsValue(3)) { triplesCount++; }
    }
  }

  public int getChecksum() {
    parseFile();
    findDoublesAndTriples();
    return doublesCount * triplesCount; 
  }

  public static void main(String[] args) {
    Puzzle2 puz = new Puzzle2();
    int checksum = puz.getChecksum();
    System.out.println(checksum);
  }
}
