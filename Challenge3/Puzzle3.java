import java.util.*;
import java.io.*;

public class Puzzle3 {
  private ArrayList<int[]> visitedCoords = new ArrayList<int[]>();
  private ArrayList<Claim> claims = new ArrayList<Claim>();
  public final static int HEIGHT = 1100;
  public final static int WIDTH = 1100;
  private int[][] fabric = new int[Puzzle3.HEIGHT][Puzzle3.WIDTH];
  private int overlap = 0;

  public int getOverlap() {
    return this.overlap;
  }

  public void findAllOverlap() {
    parseInput();
    addClaimsToFabric();
    findOverlap();
  }

  private void parseInput() {
    try {
      FileInputStream fstream = new FileInputStream("input.txt");
      BufferedReader reader = new BufferedReader(new InputStreamReader(fstream));
      String line;
      while((line = reader.readLine()) != null) {
        parseLine(line);
      }
    } catch (FileNotFoundException e) {
      System.err.println("Error: " + e.getMessage());
    } catch (IOException e) {
      System.err.println("Error: " + e.getMessage());
    }
  }

  private void parseLine(String line) {
    Claim claim = new Claim(line);
    claims.add(claim);
  }

  private void addClaimsToFabric() {
    for(Claim claim : claims) {
      for(int i = 1; i <= claim.getHeight(); i++) {
        int column = claim.getVerticalOffset() + i;
        
        for(int j = 1; j <= claim.getWidth(); j++) {
          int row = claim.getHorizontalOffset() + j;
          if(fabric[row][column] == 0) {
            fabric[row][column] = 1;
          } else {
            fabric[row][column] = 2;
          }
        }
      }
    }
  }

  private void findOverlap() {
    for(int i = 0; i < Puzzle3.HEIGHT; i++) {
      for(int j = 0; j < Puzzle3.WIDTH; j++) {
        if(fabric[i][j] == 2) {
          overlap += 1;
        }
      }
    }
//    this.overlap = visitedCoords.size();
  }

  public static void main(String[] args) {
    Puzzle3 puz = new Puzzle3(); 
    puz.findAllOverlap();
    System.out.println(puz.getOverlap());
  } 
}
