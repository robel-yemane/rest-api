# TODO for Go Album Store Project

## Project Structure
- [ ] Consider organizing your code further as the project grows, perhaps by feature or by layer (e.g., `handlers`, `models`, `db`).

## Error Handling
- [ ] Revisit usage of `log.Fatal(err)` for a more graceful error-handling technique, especially in a long-running service.

## Database Connection
- [ ] Encapsulate the database connection better, possibly by moving it to a separate package or function that can be called from `main()`.

## Environment Variables
- [ ] Explore using a configuration management library for better handling of environment variables.

## Echo Framework and API Endpoints
- [ ] Further modularise API endpoints (`getAlbums`, `getAlbumByID`, `postAlbums`) as the application grows.

## Code Comments
- [ ] Remove commented-out code for cleaner commits and maintainability.

## Podman Compose File
- [ ] Continue using Podman for containerization, and explore more advanced features that could be beneficial for deployment.

