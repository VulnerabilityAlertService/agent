#!/bin/sh

is_arm() {
  ARC=`arch | grep arm`
  test "`arch | grep arm`" = ""
  echo $?
}

keyword() {
  if [ "`getconf LONG_BIT`" = 64 ] ; then
    if [ "`is_arm`" = 0 ] ; then
      echo "64bit"
    else
      echo "ARM64"
    fi
  else
    if [ "`is_arm`" = 0 ] ; then
      echo "32bit"
    else
      echo "ARMv6"
    fi
  fi
}

has_command() {
  type $1 > /dev/null 2>&1
}

execute_command() {
  if has_command "curl"; then
    echo "curl -sSL"
  elif has_command "wget"; then
    echo "wget -q -O -"
  else
    echo "We need commands, 'tar' and 'curl' or 'wget'!" >&2
    exit 1
  fi  
}

KEYWORD=`keyword`
CMD=`execute_command`
ARCHIVE_URL=`${CMD} https://api.github.com/repos/VulnerabilityAlertService/vas/releases/latest  | grep "browser_download_url.*gz"| grep "${KEYWORD}" | cut -d : -f 2,3 | tr -d \"`

echo "Download ->${ARCHIVE_URL}"
${CMD} $ARCHIVE_URL | tar zx

if [ ! $? = 0 ] ; then
  echo "Could not get the 'vas' command!" >&2
  exit 1
fi

chmod +x ./vas

./vas -t $TOKEN
if [ ! $? = 0 ] ; then
  echo "vas failed!" >&2
  exit 1
fi

rm -rf ./vas
