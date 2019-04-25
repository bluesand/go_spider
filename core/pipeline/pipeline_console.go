package pipeline

import (
    "github.com/bluesand/go_spider/core/common/com_interfaces"
    "github.com/bluesand/go_spider/core/common/page_items"

    log "github.com/sirupsen/logrus"
)

type PipelineConsole struct {
}

func NewPipelineConsole() *PipelineConsole {
    return &PipelineConsole{}
}

func (this *PipelineConsole) Process(items *page_items.PageItems, t com_interfaces.Task) {
    log.Debug("----------------------------------------------------------------------------------------------")
    log.Debug("Crawled url :\t" + items.GetRequest().GetUrl() + "\n")
    log.Debug("Crawled result : ")
    for key, value := range items.GetAll() {
        log.Debug(key + "\t:\t" + value)
    }
}
