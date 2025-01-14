module.exports = {
  repositoryUrl: "https://github.com/insanilah/go-crud-postgres.git",
  branches: [
    { name: 'main' },         // Branch utama untuk rilis stabil
    { name: 'dev', prerelease: true }, // Branch development untuk pre-release
    { name: 'staging', prerelease: true }, // Branch staging untuk pre-release
  ],
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
