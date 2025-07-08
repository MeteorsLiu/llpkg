package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspEtmChannelT struct {
	Unused [8]uint8
}
type EspEtmChannelHandleT *EspEtmChannelT

type EspEtmEventT struct {
	EventId    c.Uint32T
	TrigPeriph EtmTriggerPeripheralT
	Del        c.Pointer
}
type EspEtmEventHandleT *EspEtmEventT

type EspEtmTaskT struct {
	TaskId     c.Uint32T
	TrigPeriph EtmTriggerPeripheralT
	Del        c.Pointer
}
type EspEtmTaskHandleT *EspEtmTaskT

/**
 * @brief ETM channel configuration
 */

type EspEtmChannelConfigT struct {
	Flags EtmChanFlags
}

type EtmChanFlags struct {
	Unused [8]uint8
}

/**
 * @brief Allocate an ETM channel
 *
 * @note The channel can later be freed by `esp_etm_del_channel`
 *
 * @param[in] config ETM channel configuration
 * @param[out] ret_chan Returned ETM channel handle
 * @return
 *      - ESP_OK: Allocate ETM channel successfully
 *      - ESP_ERR_INVALID_ARG: Allocate ETM channel failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Allocate ETM channel failed because of out of memory
 *      - ESP_ERR_NOT_FOUND: Allocate ETM channel failed because all channels are used up and no more free one
 *      - ESP_FAIL: Allocate ETM channel failed because of other reasons
 */
// llgo:link (*EspEtmChannelConfigT).EspEtmNewChannel C.esp_etm_new_channel
func (recv_ *EspEtmChannelConfigT) EspEtmNewChannel(ret_chan *EspEtmChannelHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete an ETM channel
 *
 * @param[in] chan ETM channel handle that created by `esp_etm_new_channel`
 * @return
 *      - ESP_OK: Delete ETM channel successfully
 *      - ESP_ERR_INVALID_ARG: Delete ETM channel failed because of invalid argument
 *      - ESP_FAIL: Delete ETM channel failed because of other reasons
 */
//go:linkname EspEtmDelChannel C.esp_etm_del_channel
func EspEtmDelChannel(chan_ EspEtmChannelHandleT) EspErrT

/**
 * @brief Enable ETM channel
 *
 * @note This function will transit the channel state from init to enable.
 *
 * @param[in] chan ETM channel handle that created by `esp_etm_new_channel`
 * @return
 *      - ESP_OK: Enable ETM channel successfully
 *      - ESP_ERR_INVALID_ARG: Enable ETM channel failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Enable ETM channel failed because the channel has been enabled already
 *      - ESP_FAIL: Enable ETM channel failed because of other reasons
 */
//go:linkname EspEtmChannelEnable C.esp_etm_channel_enable
func EspEtmChannelEnable(chan_ EspEtmChannelHandleT) EspErrT

/**
 * @brief Disable ETM channel
 *
 * @note This function will transit the channel state from enable to init.
 *
 * @param[in] chan ETM channel handle that created by `esp_etm_new_channel`
 * @return
 *      - ESP_OK: Disable ETM channel successfully
 *      - ESP_ERR_INVALID_ARG: Disable ETM channel failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Disable ETM channel failed because the channel is not enabled yet
 *      - ESP_FAIL: Disable ETM channel failed because of other reasons
 */
//go:linkname EspEtmChannelDisable C.esp_etm_channel_disable
func EspEtmChannelDisable(chan_ EspEtmChannelHandleT) EspErrT

/**
 * @brief Connect an ETM event to an ETM task via a previously allocated ETM channel
 *
 * @note Setting the ETM event/task handle to NULL means to disconnect the channel from any event/task
 *
 * @param[in] chan ETM channel handle that created by `esp_etm_new_channel`
 * @param[in] event ETM event handle obtained from a driver/peripheral, e.g. `xxx_new_etm_event`
 * @param[in] task ETM task handle obtained from a driver/peripheral, e.g. `xxx_new_etm_task`
 * @return
 *      - ESP_OK: Connect ETM event and task to the channel successfully
 *      - ESP_ERR_INVALID_ARG: Connect ETM event and task to the channel failed because of invalid argument
 *      - ESP_FAIL: Connect ETM event and task to the channel failed because of other reasons
 */
//go:linkname EspEtmChannelConnect C.esp_etm_channel_connect
func EspEtmChannelConnect(chan_ EspEtmChannelHandleT, event EspEtmEventHandleT, task EspEtmTaskHandleT) EspErrT

/**
 * @brief Delete ETM event
 *
 * @note Although the ETM event comes from various peripherals, we provide the same user API to delete the event handle seamlessly.
 *
 * @param[in] event ETM event handle obtained from a driver/peripheral, e.g. `xxx_new_etm_event`
 * @return
 *      - ESP_OK: Delete ETM event successfully
 *      - ESP_ERR_INVALID_ARG: Delete ETM event failed because of invalid argument
 *      - ESP_FAIL: Delete ETM event failed because of other reasons
 */
//go:linkname EspEtmDelEvent C.esp_etm_del_event
func EspEtmDelEvent(event EspEtmEventHandleT) EspErrT

/**
 * @brief Delete ETM task
 *
 * @note Although the ETM task comes from various peripherals, we provide the same user API to delete the task handle seamlessly.
 *
 * @param[in] task ETM task handle obtained from a driver/peripheral, e.g. `xxx_new_etm_task`
 * @return
 *      - ESP_OK: Delete ETM task successfully
 *      - ESP_ERR_INVALID_ARG: Delete ETM task failed because of invalid argument
 *      - ESP_FAIL: Delete ETM task failed because of other reasons
 */
//go:linkname EspEtmDelTask C.esp_etm_del_task
func EspEtmDelTask(task EspEtmTaskHandleT) EspErrT

/**
 * @brief Dump ETM channel usages to the given IO stream
 *
 * @param[in] out_stream IO stream (e.g. stdout)
 * @return
 *      - ESP_OK: Dump ETM channel usages successfully
 *      - ESP_ERR_INVALID_ARG: Dump ETM channel usages failed because of invalid argument
 *      - ESP_FAIL: Dump ETM channel usages failed because of other reasons
 */
//go:linkname EspEtmDump C.esp_etm_dump
func EspEtmDump(out_stream *c.FILE) EspErrT
