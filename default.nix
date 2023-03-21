
with (import <nixpkgs> {});

stdenv.mkDerivation {


  name = "terragrunt-backend-simple";
  buildInputs = [
    terraform
    kubectl
    kind
    kubernetes-helm-wrapped
    hubble
    cilium-cli
    vault-bin
    jq
  ];
}
