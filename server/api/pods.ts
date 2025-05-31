import { coreV1Api } from '~/server/utils/kubernetes'

export default defineEventHandler(async () => {
    const config = useRuntimeConfig()
    const res = await coreV1Api.listPodForAllNamespaces({
        labelSelector: config.kubernetes.labelSelector,
    })
    return res
})