import * as k8s from '@kubernetes/client-node'

let coreV1Api: k8s.CoreV1Api
let appsV1Api: k8s.AppsV1Api
let kubeConfig: k8s.KubeConfig

export function setKubernetesClients(kc: k8s.KubeConfig) {
    kubeConfig = kc
    coreV1Api = kc.makeApiClient(k8s.CoreV1Api)
    appsV1Api = kc.makeApiClient(k8s.AppsV1Api)
    attach = new k8s.Att
}

export { coreV1Api, appsV1Api, kubeConfig }
