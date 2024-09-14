#!/bin/bash

red='\033[0;31m'
green='\033[0;32m'
yellow='\033[0;33m'
plain='\033[0m'

cur_dir=$(pwd)

# check root
[[ $EUID -ne 0 ]] && echo -e "${red}Fatal error: ${plain} Please run this script with root privilege \n " && exit 1

# Check OS and set release variable
if [[ -f /etc/os-release ]]; then
  source /etc/os-release
  release=$ID
elif [[ -f /usr/lib/os-release ]]; then
  source /usr/lib/os-release
  release=$ID
else
  echo "Failed to check the system OS, please contact the author!" >&2
  exit 1
fi
echo "The OS release is: $release"

arch() {
  case "$(uname -m)" in
  x86_64 | x64 | amd64) echo 'amd64' ;;
  i*86 | x86) echo '386' ;;
  armv8* | armv8 | arm64 | aarch64) echo 'arm64' ;;
  armv7* | armv7 | arm) echo 'armv7' ;;
  armv6* | armv6) echo 'armv6' ;;
  armv5* | armv5) echo 'armv5' ;;
  s390x) echo 's390x' ;;
  *) echo -e "${green}Unsupported CPU architecture! ${plain}" && rm -f install.sh && exit 1 ;;
  esac
}

echo "arch: $(arch)"

os_version=""
os_version=$(grep -i version_id /etc/os-release | cut -d \" -f2 | cut -d . -f1)

if [[ "${release}" == "arch" ]]; then
  echo "Your OS is Arch Linux"
elif [[ "${release}" == "parch" ]]; then
  echo "Your OS is Parch linux"
elif [[ "${release}" == "manjaro" ]]; then
  echo "Your OS is Manjaro"
elif [[ "${release}" == "armbian" ]]; then
  echo "Your OS is Armbian"
elif [[ "${release}" == "opensuse-tumbleweed" ]]; then
  echo "Your OS is OpenSUSE Tumbleweed"
elif [[ "${release}" == "centos" ]]; then
  if [[ ${os_version} -lt 8 ]]; then
    echo -e "${red} Please use CentOS 8 or higher ${plain}\n" && exit 1
  fi
elif [[ "${release}" == "ubuntu" ]]; then
  if [[ ${os_version} -lt 20 ]]; then
    echo -e "${red} Please use Ubuntu 20 or higher version!${plain}\n" && exit 1
  fi
elif [[ "${release}" == "fedora" ]]; then
  if [[ ${os_version} -lt 36 ]]; then
    echo -e "${red} Please use Fedora 36 or higher version!${plain}\n" && exit 1
  fi
elif [[ "${release}" == "debian" ]]; then
  if [[ ${os_version} -lt 11 ]]; then
    echo -e "${red} Please use Debian 11 or higher ${plain}\n" && exit 1
  fi
elif [[ "${release}" == "almalinux" ]]; then
  if [[ ${os_version} -lt 9 ]]; then
    echo -e "${red} Please use AlmaLinux 9 or higher ${plain}\n" && exit 1
  fi
elif [[ "${release}" == "rocky" ]]; then
  if [[ ${os_version} -lt 9 ]]; then
    echo -e "${red} Please use Rocky Linux 9 or higher ${plain}\n" && exit 1
  fi
elif [[ "${release}" == "oracle" ]]; then
  if [[ ${os_version} -lt 8 ]]; then
    echo -e "${red} Please use Oracle Linux 8 or higher ${plain}\n" && exit 1
  fi
else
  echo -e "${red}Your operating system is not supported by this script.${plain}\n"
  echo "Please ensure you are using one of the following supported operating systems:"
  echo "- Ubuntu 20.04+"
  echo "- Debian 11+"
  echo "- CentOS 8+"
  echo "- Fedora 36+"
  echo "- Arch Linux"
  echo "- Parch Linux"
  echo "- Manjaro"
  echo "- Armbian"
  echo "- AlmaLinux 9+"
  echo "- Rocky Linux 9+"
  echo "- Oracle Linux 8+"
  echo "- OpenSUSE Tumbleweed"
  exit 1

fi
