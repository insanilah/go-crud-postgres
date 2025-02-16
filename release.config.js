module.exports = {
  repositoryUrl: "https://github.com/insanilah/go-crud-postgres.git",
  branches: [
    { name: "main" }, // Branch utama (rilis stabil)
    { name: "staging", prerelease: "beta" }, // Branch staging → vX.Y.Z-beta.N
    { name: "dev", prerelease: "dev" }, // Branch dev → vX.Y.Z-dev.N
  ],
  plugins: [
    "@semantic-release/commit-analyzer",
    "@semantic-release/release-notes-generator",
    "@semantic-release/changelog",
    [
      "@semantic-release/github",
      {
        assets: [{ path: "dist/*.zip", label: "Build Package" }],
      },
    ],
  ],
};
