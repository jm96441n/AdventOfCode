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
        int row = claim.getVerticalOffset() + i;
        
        for(int j = 1; j <= claim.getWidth(); j++) {
          int column = claim.getHorizontalOffset() + j;
          if(fabric[row][column] == 0) {
            fabric[row][column] = 1;
          } else {
            if(fabric[row][column] == 1) {
              int[] coords = { row, column };
              visitedCoords.add(coords);
            }
            fabric[row][column] = 2;
          }
        }
      }
    }
  }

  private void findOverlap() {
    this.overlap = visitedCoords.size();
  }

  public String findNoOverlap() {
    String claimId = "";
    for(Claim claim : claims) {
      boolean foundClaim = true;
      for(int i = 1; i <= claim.getHeight(); i++) {
        int row = claim.getVerticalOffset() + i;

        for(int j = 1; j <= claim.getWidth(); j++) {
          int column = claim.getHorizontalOffset() + j;
          if(fabric[row][column] == 2) {
            foundClaim = false;
            break;
          }
        }

        if(foundClaim == false) { break; }
      }

      if(foundClaim == true) {
        claimId = claim.getClaimNumber();
        break;
      }
    }
    // will never hit this
    return claimId;
  }

  public static void main(String[] args) {
    Puzzle3 puz = new Puzzle3(); 
    puz.findAllOverlap();
    System.out.println(puz.getOverlap());
    String idNoOverlap = puz.findNoOverlap();
    System.out.println(idNoOverlap);
  } 
}
