module.exports = {
  repositoryUrl: "https://github.com/insanilah/go-crud-postgres.git",
  branches: ["main", "staging", "dev"],
  plugins: [
    "@semantic-release/commit-analyzer",
    "@semantic-release/release-notes-generator",
    "@semantic-release/changelog",
    [
      "@semantic-release/github",
      {
        assets: [
          { path: "dist/*.zip", label: "Build Package" }
        ]
      }
    ]
  ]
};