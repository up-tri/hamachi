
main() {
  export GOOGLE_APPLICATION_CREDENTIALS=${PROJ_ROOT}/keys/firebase-admin.json
  cd ${PROJ_ROOT}/emulators/firestore && firebase emulators:start
}

SCRIPT_DIR="$(cd "$(dirname "${0}")" && echo "${PWD}")"
PROJ_ROOT="${SCRIPT_DIR}/.."
main "$@"
