main() {
  go test -v ${PROJ_ROOT}/...
}

SCRIPT_DIR="$(cd "$(dirname "${0}")" && echo "${PWD}")"
PROJ_ROOT="${SCRIPT_DIR}/.."
main "$@"
