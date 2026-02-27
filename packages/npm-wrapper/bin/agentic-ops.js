#!/usr/bin/env node

const { spawn } = require('child_process');
const path = require('path');
const fs = require('fs');
const os = require('os');

// Determine the binary name based on platform and architecture
function getBinaryName() {
  const platform = os.platform();
  const arch = os.arch();

  let osName;
  switch (platform) {
    case 'win32':
      osName = 'windows';
      break;
    case 'darwin':
      osName = 'darwin';
      break;
    case 'linux':
      osName = 'linux';
      break;
    default:
      console.error(`Unsupported platform: ${platform}`);
      process.exit(1);
  }

  let archName;
  switch (arch) {
    case 'x64':
    case 'amd64':
      archName = 'amd64';
      break;
    case 'arm64':
    case 'aarch64':
      archName = 'arm64';
      break;
    default:
      console.error(`Unsupported architecture: ${arch}`);
      process.exit(1);
  }

  const ext = platform === 'win32' ? '.exe' : '';
  return `agentic-ops-${osName}-${archName}${ext}`;
}

// Find the binary
function findBinary() {
  const binaryName = getBinaryName();
  const binDir = path.join(__dirname);
  const binaryPath = path.join(binDir, binaryName);

  if (fs.existsSync(binaryPath)) {
    return binaryPath;
  }

  // Also check parent directory in case of different install layouts
  const parentBinDir = path.join(__dirname, '..', 'bin');
  const parentBinaryPath = path.join(parentBinDir, binaryName);
  if (fs.existsSync(parentBinaryPath)) {
    return parentBinaryPath;
  }

  console.error(`Binary not found: ${binaryName}`);
  console.error(`Looked in: ${binDir}, ${parentBinDir}`);
  console.error('Please reinstall the package: npm install -g agentic-ops');
  process.exit(1);
}

// Execute the binary with all arguments passed through
const binary = findBinary();
const args = process.argv.slice(2);

const child = spawn(binary, args, {
  stdio: 'inherit',
  shell: false
});

child.on('error', (err) => {
  console.error(`Failed to execute binary: ${err.message}`);
  process.exit(1);
});

child.on('close', (code) => {
  process.exit(code || 0);
});
