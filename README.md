# Groupie-Tracker

Groupie-Tracker is a project written in Golang aimed at receiving a given [API](https://groupietrackers.herokuapp.com/api) and manipulating the contained data to create a website displaying the information. The provided APIs contain music group names, along with their albums, artists, creation dates, first album dates, and concert dates.
<p align="center">
  <img src="picturesReadme/home.png" alt="home">
</p>

## Features

- Display a list of music groups on the homepage.
- Click on a group's image to navigate to its dedicated : 
   - Members.
   - Creation Date.
   - First Album.
   - Concert Dates.
   <p align="center">
   <img src="picturesReadme/artist_page.png" alt="artist_page" width="300">
   </p>
- Click on one of them to view displaying detailed information :
   <p align="center">
   <img src="picturesReadme/artist_detailed.png" alt="artist_detailed" width="300">
   </p>
- Search functionality allows users to search for groups, artists, album names, and concert venues.
- Filters enable users to refine the list of displayed groups on the homepage.
- Styled with CSS inspired by JoJo's Bizarre Adventure.

## Installation

To run Groupie-Tracker locally, follow these steps:

1. Clone this repository to your local machine.
   ```
   git clone https://github.com/yourusername/groupie-tracker.git
   ```
2. Navigate to the project directory.
   ```
   cd groupie-tracker
   ```
3. Build the project.
   ```
   go build cmd/main.go
   ```
4. Run the executable.
   ```
   ./groupie-tracker
   ```
5. if you don't want build, you could run the project.
   ```
   go run cmd/main.go
   ```
6. Access the application in your web browser at `http://localhost:8080` or click in your terminal on the link.

## Usage

Upon running Groupie-Tracker, you can access the homepage where a list of music groups is displayed. You can click on any group's image to view detailed information about that group on a separate page. The search bar allows you to search for specific groups, artists, albums, or concert venues. Additionally, filters are available to narrow down the displayed groups based on certain criteria.

## Co-developers

- Delestre Thomas
- Fouquemberg Axelle
- Lovergne Raphael