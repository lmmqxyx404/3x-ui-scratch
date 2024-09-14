#!/bin/bash

# define the output log color
red='\033[0;31m'
green='\033[0;32m'
yellow='\033[0;33m'
plain='\033[0m'

#Add some basic function here
function LOGD() {
  echo -e "${yellow}[DEG] $* ${plain}"
}

function LOGE() {
  echo -e "${red}[ERR] $* ${plain}"
}

function LOGI() {
  echo -e "${green}[INF] $* ${plain}"
}

# check root 检测 root 权限，必须使用 root 权限执行 shell
[[ $EUID -ne 0 ]] && LOGE "ERROR: You must be root to run this script! \n" && exit 1

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

os_version=""
os_version=$(grep -i version_id /etc/os-release | cut -d \" -f2 | cut -d . -f1)
echo "The os_version is: $os_version"

function check_target_os_version() {
  local release=$1
  local os_version=$2
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
}

check_target_os_version "$release" "$os_version"

show_menu() {
  echo -e "
  ${green}3X-UI Panel Management Script${plain}
  ${green}0.${plain} Exit Script
————————————————
  ${green}1.${plain} Install
  ${green}2.${plain} Update
  ${green}3.${plain} Update Menu
  ${green}4.${plain} Custom Version
  ${green}5.${plain} Uninstall
————————————————
  ${green}6.${plain} Reset Username & Password & Secret Token
  ${green}7.${plain} Reset Web Base Path
  ${green}8.${plain} Reset Settings
  ${green}9.${plain} Change Port
  ${green}10.${plain} View Current Settings
————————————————
  ${green}11.${plain} Start
  ${green}12.${plain} Stop
  ${green}13.${plain} Restart
  ${green}14.${plain} Check Status
  ${green}15.${plain} Check Logs
————————————————
  ${green}16.${plain} Enable Autostart
  ${green}17.${plain} Disable Autostart
————————————————
  ${green}18.${plain} SSL Certificate Management
  ${green}19.${plain} Cloudflare SSL Certificate
  ${green}20.${plain} IP Limit Management
  ${green}21.${plain} Firewall Management
————————————————
  ${green}22.${plain} Enable BBR 
  ${green}23.${plain} Update Geo Files
  ${green}24.${plain} Speedtest by Ookla
"

  echo && read -p "Please enter your selection [0-24]: " num
}

# 定义 main 函数
main() {
  echo "函数接收到 $# 个参数"

  if [[ $# > 0 ]]; then

  else
    show_menu
  fi

}

# 调用 main 函数并将所有传递给脚本的参数传递给 main
main "$@"
