import * as k8s from '@kubernetes/client-node'
import { setKubernetesClients } from '~/server/utils/kubernetes'

export default defineNitroPlugin(() => {
    const kc = new k8s.KubeConfig()

    try {
        kc.loadFromDefault()
        setKubernetesClients(kc)
        console.log('✅ Kubernetes client initialized')
    } catch (err) {
        console.error('❌ Failed to load kubeconfig:', err)
    }
})
