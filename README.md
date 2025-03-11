# CI-CD-Security-Scanner

The **CI/CD Security Scanner** is a command-line tool built in Go that scans CI/CD pipeline configurations (GitHub Actions, GitLab CI/CD) for security misconfigurations and policy violations. It helps DevOps and security teams enforce best practices and reduce the attack surface of their pipelines.

## Features
- Parses CI/CD YAML configurations (GitHub Actions, GitLab CI/CD).
- Identifies insecure practices such as:
  - Running jobs as `root`.
  - Using `latest` tags in Docker images.
  - Hardcoded credentials.
- Uses **Open Policy Agent (OPA)** for policy enforcement.
- Generates a security report.
- Can be integrated into CI/CD pipelines to fail builds on violations.

## Installation
Ensure you have **Go 1.20+** installed on your system.

Clone the repository:
```sh
git clone https://github.com/iamdanielokoye/CI-CD-Security-Scanner.git
cd CI-CD-Security-Scanner
```

Build the tool:
```sh
go build -o scanner cmd/main.go
```

## Usage
Run the scanner against a CI/CD pipeline YAML file:
```sh
./scanner examples/insecure-github.yml
```

This will analyze the provided YAML configuration and output detected security issues.

## Integration with CI/CD Pipelines
You can add the scanner to your CI/CD pipeline for automated security checks.

### GitHub Actions Example
Modify your **workflow-test.yml** to include:
```yaml
- name: Run Security Scanner
  run: go run cmd/main.go examples/insecure-github.yml
```

### GitLab CI/CD Example
Add a security scan job:
```yaml
security_scan:
  stage: test
  image: golang:1.20
  script:
    - go run cmd/main.go examples/secure-gitlab.yml
```

## Testing
Run the tests using:
```sh
go test ./test/... --cover
```

## License
This project is licensed under the **MIT License**. See [LICENSE](LICENSE) for details.

