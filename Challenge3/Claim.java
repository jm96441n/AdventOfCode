public class Claim {
  private String claimNumber;
  private int horizontalOffset;
  private int verticalOffset;
  private int height;
  private int width;

  public Claim(String line) {
    parseClaim(line);
  }

  private void parseClaim(String claim) {
    String[] splitLine = claim.split(" ");
    this.setClaimNumber(splitLine[0]);
    this.setHorizontalOffset(splitLine[2]);
    this.setVerticalOffset(splitLine[2]);
    this.setWidth(splitLine[3]);
    this.setHeight(splitLine[3]);
  }

  private void setClaimNumber(String claimNumber) {
    this.claimNumber = claimNumber;
  }

  private void setHorizontalOffset(String offsets) {
    String horizontalString = offsets.split(",")[0];
    int horizontal = Integer.parseInt(horizontalString);
    this.horizontalOffset = horizontal;
  }
  
  private void setVerticalOffset(String offsets) {
    String verticalString = offsets.split(",")[1].split(":")[0];
    int vertical = Integer.parseInt(verticalString);
    this.verticalOffset = vertical;
  }

  private void setWidth(String area) {
    String widthString = area.split("x")[0];
    int width = Integer.parseInt(widthString);
    this.width = width;
  }

  private void setHeight(String area) {
    String heightString = area.split("x")[1];
    int height = Integer.parseInt(heightString);
    this.height = height;
  }

  public String getClaimNumber() {
    return this.claimNumber;
  }

  public int getHorizontalOffset() {
    return this.horizontalOffset;
  }

  public int getVerticalOffset() {
    return this.verticalOffset;
  }
 
  public int getWidth() {
    return this.width; 
  }

  public int getHeight() {
    return this.height;
  }

  public String toString() {
    return getClaimNumber() + " @ " + getHorizontalOffset() + "," + getVerticalOffset() + ": " + getWidth() + "x" + getHeight();
  }
}
